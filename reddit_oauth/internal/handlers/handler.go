package handlers

import (
	"errors"
	"net/http"

	"github.com/SantiagoBedoya/reddit_oauth/internal/dto"
	"github.com/SantiagoBedoya/reddit_oauth/internal/models"
	"github.com/SantiagoBedoya/reddit_oauth/internal/service"
	"github.com/SantiagoBedoya/reddit_oauth/pkg/validate"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Handler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) Handler {
	return &handler{service}
}

func (h handler) Register(c *fiber.Ctx) error {
	var registeDto dto.RegisterDto
	if err := c.BodyParser(&registeDto); err != nil {
		httpErr := validate.HTTPError{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
		return c.Status(httpErr.StatusCode).JSON(httpErr)
	}
	err := validate.Validate(registeDto)
	if err != nil {
		return c.Status(err.StatusCode).JSON(err)
	}
	if err := h.service.Register(&models.User{
		Username: registeDto.Username,
		Email:    registeDto.Email,
		Password: registeDto.Password,
	}); err != nil {
		logrus.Error(err)
		if errors.Is(err, service.ErrEmailInUse) {
			httpErr := validate.HTTPError{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
			}
			return c.Status(httpErr.StatusCode).JSON(httpErr)
		}
		httpErr := validate.HTTPError{
			StatusCode: http.StatusInternalServerError,
			Message:    "something went wrong",
		}
		return c.Status(httpErr.StatusCode).JSON(httpErr)
	}
	return c.SendStatus(http.StatusCreated)
}

func (h handler) Login(c *fiber.Ctx) error {
	var loginDto dto.LoginDto
	if err := c.BodyParser(&loginDto); err != nil {
		httpErr := validate.HTTPError{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
		return c.Status(httpErr.StatusCode).JSON(httpErr)
	}

	if err := validate.Validate(loginDto); err != nil {
		return c.Status(err.StatusCode).JSON(err)
	}
	token, err := h.service.Login(&models.User{
		Email:    loginDto.Email,
		Password: loginDto.Password,
	})
	if err != nil {
		logrus.Error(err)
		if errors.Is(err, service.ErrInvalidCredentials) {
			httpErr := validate.HTTPError{
				StatusCode: http.StatusUnauthorized,
				Message:    err.Error(),
			}
			return c.Status(httpErr.StatusCode).JSON(httpErr)
		}
		httpErr := validate.HTTPError{
			StatusCode: http.StatusInternalServerError,
			Message:    "something went wrong",
		}
		return c.Status(httpErr.StatusCode).JSON(httpErr)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"access_token": token})
}
