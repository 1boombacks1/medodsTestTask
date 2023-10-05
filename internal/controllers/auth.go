package controllers

import (
	"d0c/TestTaskBackDev/internal/services"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type response struct {
	AuthToken string `json:"auth_token"`
}

type authRoutes struct {
	sessionService services.Session
}

func newAuthRoutes(g fiber.Router, sessionService services.Session) {
	r := authRoutes{sessionService: sessionService}
	g.Get("/generateTokens", r.generateTokens)
	g.Get("/refresh", r.refreshToken)
}

// @Summary GenerateTokens
// @Description Generate tokens by guid in query
// @Tags auth
// @Produce json
// @Param guid query string true "guid"
// @Success 200 {object} response
// @Failure 401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /auth/generateTokens [get]
func (r *authRoutes) generateTokens(c *fiber.Ctx) error {
	guid := c.Query("guid")
	if len(guid) != 32 {
		newErrorResponse(c, fiber.StatusUnauthorized, "INVALID_GUID")
		return errGuidNotValid
	}

	tokens, err := r.sessionService.GenerateTokens(c.Context(), guid)
	if err != nil {
		newErrorResponse(c, fiber.StatusInternalServerError, "INTERNAL_SERVER_ERROR")
		return errInternalServerErr
	}

	c.Cookie(&fiber.Cookie{
		HTTPOnly: true,
		Name:     "RefreshToken",
		Path:     "/auth",
		Value:    tokens.RefreshToken,
		MaxAge:   tokens.RefreshTokenTTL,
	})

	return c.Status(fiber.StatusOK).JSON(response{
		AuthToken: tokens.AccessToken,
	})
}

// @Summary RefreshTokens
// @Description Refresh tokens by given valid 'refresh token' in cookies
// @Tags auth
// @Produce json
// @Success 200 {object} response
// @Failure 401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /auth/refresh [get]
func (r *authRoutes) refreshToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("RefreshToken")
	if len(refreshToken) == 0 {
		newErrorResponse(c, fiber.StatusUnauthorized, "EMPTY_REFRESH_TOKEN")
		return errRefreshTokenEmpty
	}

	fmt.Printf("RT - %s\n", refreshToken)

	tokens, err := r.sessionService.RefreshTokens(c.Context(), refreshToken)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrInvalidRefreshSession):
			newErrorResponse(c, fiber.StatusUnauthorized, "INVALID_REFRESH_TOKEN")
			return errRefreshTokenNotValid
		case errors.Is(err, services.ErrRefreshTokenExpired):
			newErrorResponse(c, fiber.StatusUnauthorized, "TOKEN_EXPIRED")
			return errTokenExpired
		default:
			newErrorResponse(c, fiber.StatusInternalServerError, "INTERNAL_SERVER_ERROR")
			return errInternalServerErr
		}
	}

	c.Cookie(&fiber.Cookie{
		HTTPOnly: true,
		Name:     "RefreshToken",
		Path:     "/auth",
		Value:    tokens.RefreshToken,
		MaxAge:   tokens.RefreshTokenTTL,
	})

	return c.Status(fiber.StatusOK).JSON(response{
		AuthToken: tokens.AccessToken,
	})
}
