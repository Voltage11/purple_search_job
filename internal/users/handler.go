package users

import (
	"fmt"
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

type UserHandler struct {
	router       fiber.Router
	customLogger zerolog.Logger
	repository *UserRepository
}

func NewHadnler(router fiber.Router, customLogger *zerolog.Logger, repository *UserRepository) {
	h := &UserHandler{
		router: router,
		customLogger: *customLogger,
		repository: repository,
	}

	vacancyGroup := h.router.Group("/user")
	vacancyGroup.Post("/", h.createUser)
}

func (h *UserHandler) createUser(c *fiber.Ctx) error {
	form := UserCreateForm{
		Name: c.FormValue("name"),
		Email: c.FormValue("email"),
		PassHash: c.FormValue("password"),
	}
	errors := validate.Validate(
		&validators.StringIsPresent{ Name: "Name", Field: form.Name, Message: "Укажите имя"},
		&validators.EmailIsPresent{ Name: "Email", Field: form.Email, Message: "Email не задан или не верный"},
		&validators.StringIsPresent{ Name: "PassHash", Field: form.PassHash, Message: "Укажите пароль"},		
	)
	var component templ.Component
	if (len(errors.Errors) > 0) {
		component = components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return tadapter.Render(c, component, http.StatusBadRequest)
	}
	email, err := h.repository.AddUser(form)
	if (err != nil) {
		h.customLogger.Error().Msg(err.Error())
		component = components.Notification("Произошла ошибка добавления пользователя", components.NotificationFail)
		return tadapter.Render(c, component, http.StatusBadRequest)
	}
	component = components.Notification(fmt.Sprintf("Пользователь создан: %s", email), components.NotificationSuccess)
	return tadapter.Render(c, component, http.StatusOK)

}