package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

var (
	errGuidNotValid         = errors.New("guid not valid")
	errRefreshTokenNotValid = errors.New("refresh token not valid")
	errRefreshTokenEmpty    = errors.New("refresh token empty")
	errTokenExpired         = errors.New("refresh token expired")
	errInternalServerErr    = errors.New("internal server error")
)

type errorResponse struct {
	StatusCode   int
	ErrorMessage string
}

func newErrorResponse(c *fiber.Ctx, errStatus int, message string) {
	c.Status(errStatus).JSON(errorResponse{
		StatusCode:   errStatus,
		ErrorMessage: message,
	})
}
