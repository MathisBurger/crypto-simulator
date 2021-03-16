///////////////////////////////////////////////////
// This file is not longer in use.
// It can be enabled via config
///////////////////////////////////////////////////

package controller

import (
	"encoding/json"
	"github.com/MathisBurger/crypto-simulator/database/actions"
	"github.com/MathisBurger/crypto-simulator/utils"
	"github.com/gofiber/fiber/v2"
)

// DEPRECATED
type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// DEPRECATED
type loginResponse struct {
	Status    bool   `json:"status"`
	Message   string `json:"message"`
	AuthToken string `json:"auth_token"`
}

// DEPRECATED
func LoginController(c *fiber.Ctx) error {

	// parsing and checking request
	obj := loginRequest{}
	err := json.Unmarshal(c.Body(), &obj)

	if err != nil {
		return c.JSON(loginResponse{
			false,
			"Invalid JSON body",
			"None",
		})
	}
	if !checkLoginRequest(obj) {
		return c.JSON(loginResponse{
			false,
			"Invalid JSON body",
			"None",
		})
	}

	// execute login
	if actions.Login(obj.Username, obj.Password) {

		// generate and set token
		token := utils.GenerateToken()
		actions.SetUserAuthToken(obj.Username, token)

		return c.JSON(loginResponse{
			true,
			"successfully logged in",
			token,
		})

	} else {
		return c.JSON(loginResponse{
			false,
			"wrong login credentials",
			"None",
		})
	}
}

// checks request
// DEPRECATED
func checkLoginRequest(obj loginRequest) bool {
	return obj.Username != "" && obj.Password != ""
}
