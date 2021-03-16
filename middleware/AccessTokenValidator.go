package middleware

import (
	"fmt"
	"github.com/MathisBurger/crypto-simulator/accesstoken"
	"github.com/gofiber/fiber/v2"
	"strings"
)

var atvalidator accesstoken.Validator

// This functions validates the JWT
// It requires the context of the called
// API endpoint, to perform
func ValidateAccessToken(c *fiber.Ctx) (bool, string) {
	atvalidator, _ = accesstoken.NewJWTManager("", "./certs/public.pem")

	authheader := c.Get("Authorization")

	if !strings.HasPrefix(authheader, "accessToken ") {
		fmt.Println("benz weiß")
		return false, ""
	}
	accessToken := authheader[12:]
	ident, err := atvalidator.Validate(accessToken)
	if err != nil {
		fmt.Println(err.Error())
		return false, ""
	}

	return true, ident
}
