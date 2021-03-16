package auth

import (
	"github.com/MathisBurger/crypto-simulator/database/actions"
	"github.com/MathisBurger/crypto-simulator/middleware"
	"github.com/gofiber/fiber/v2"
)

// This endpoint revokes an existing session based on its credentials
func RevokeSessionController(c *fiber.Ctx) error {

	if status, ident := middleware.ValidateAccessToken(c); status {

		if actions.CheckUserOwnsRefreshToken(ident, c.Cookies("refreshToken")) {

			actions.RevokeSession(ident, c.Cookies("refreshToken"))

			return c.SendStatus(200)

		} else {
			return c.SendStatus(401)
		}

	} else {
		return c.SendStatus(401)
	}
}
