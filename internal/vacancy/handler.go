package vacancy

import (
	"lesson/pkg/tadapter"
	"lesson/pkg/validator"
	"lesson/views/components"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type VacancyHandler struct {
	router       fiber.Router
	customLogger zerolog.Logger
	repository *VacancyRepository
}

func NewHadnler(router fiber.Router, customLogger *zerolog.Logger, repository *VacancyRepository) {
	h := &VacancyHandler{
		router: router,
		customLogger: *customLogger,
		repository: repository,
	}

	vacancyGroup := h.router.Group("/vacancy")
	vacancyGroup.Post("/", h.createVacancy)
}

func (h *VacancyHandler) createVacancy(c *fiber.Ctx) error {
	form := VacancyCreateForm{
		Role: c.FormValue("role"),
		Company: c.FormValue("company"),
		Type: c.FormValue("type"),
		Salary: c.FormValue("salary"),
		Location: c.FormValue("location"),
		Email: c.FormValue("email"),
	}
	errors := validate.Validate(
		&validators.StringIsPresent{ Name: "Role", Field: form.Role, Message: "Укажите должность"},
		&validators.StringIsPresent{ Name: "company", Field: form.Company, Message: "Укажите компанию"},
		&validators.StringIsPresent{ Name: "type", Field: form.Type, Message: "Укажите сферу компании"},
		&validators.StringIsPresent{ Name: "salary", Field: form.Salary, Message: "Укажите зп"},
		&validators.StringIsPresent{ Name: "location", Field: form.Location, Message: "Укажите расположение"},
		&validators.EmailIsPresent{ Name: "Email", Field: form.Email, Message: "Email не задан или не верный"},
	)
	var component templ.Component
	if (len(errors.Errors) > 0) {
		component = components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return tadapter.Render(c, component, http.StatusBadRequest)
	}
	err := h.repository.addVacancy(form)
	if (err != nil) {
		h.customLogger.Error().Msg(err.Error())
		component = components.Notification("Произошла ошибка добавления вакансии", components.NotificationFail)
		return tadapter.Render(c, component, http.StatusBadRequest)
	}
	component = components.Notification("Вакансия создана", components.NotificationSuccess)
	return tadapter.Render(c, component, http.StatusOK)

}