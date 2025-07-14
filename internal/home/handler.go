package home

import (
	"lesson/pkg/tadapter"
	"lesson/views"
	"math"
	"net/http"

	"lesson/internal/vacancy"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router 		 fiber.Router
	customLogger zerolog.Logger
	repository *vacancy.VacancyRepository
}

type User struct {
	Id int
	Name string
}

func NewHadnler(router fiber.Router, customLogger *zerolog.Logger, repository *vacancy.VacancyRepository) {
	h := HomeHandler{
		router: router,
		customLogger: *customLogger,
		repository: repository,
	}
	//api := h.router.Group("/api")

	h.router.Get("/", h.home)
	h.router.Get("/404", h.error)
}

func (h *HomeHandler) home (c *fiber.Ctx) error {
	PAGE_ITEMS := 2
	page := c.QueryInt("page", 1)
	count := h.repository.CountAll()

	vacancies, err := h.repository.GetAll(PAGE_ITEMS, (page - 1)*PAGE_ITEMS)
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
		return c.SendStatus(500)
	}
	

	component := views.Main(vacancies,int(math.Ceil(float64(count / PAGE_ITEMS))), page)
	
	return tadapter.Render(c, component, http.StatusOK)
}


func (h *HomeHandler) error (c *fiber.Ctx) error {
	return fiber.NewError(500, "Не указали параметры")
}