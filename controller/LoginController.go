package controller

import (
	"encoding/json"
	"github.com/MathisBurger/crypto-simulator/database/actions"
	"github.com/MathisBurger/crypto-simulator/utils"
	"github.com/gofiber/fiber/v2"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	Status    bool   `json:"status"`
	Message   string `json:"message"`
	AuthToken string `json:"auth_token"`
}

func LoginController(c *fiber.Ctx) error {
	raw := string(c.Body())
	obj := loginRequest{}
	err := json.Unmarshal([]byte(raw), &obj)
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
	if actions.Login(obj.Username, obj.Password) {
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

func checkLoginRequest(obj loginRequest) bool {
	return obj.Username != "" && obj.Password != ""
}