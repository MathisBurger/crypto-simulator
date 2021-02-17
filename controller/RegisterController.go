package controller

import (
	"encoding/json"
	"github.com/MathisBurger/crypto-simulator/database/actions"
	"github.com/gofiber/fiber/v2"
)

type registerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type registerResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func RegisterController(c *fiber.Ctx) error {
	raw := string(c.Body())
	obj := registerRequest{}
	err := json.Unmarshal([]byte(raw), &obj)
	if err != nil {
		return c.JSON(registerResponse{
			false,
			"Invalid JSON body",
		})
	}
	if !checkRegisterRequest(obj) {
		return c.JSON(registerResponse{
			false,
			"Invalid JSON body",
		})
	}
	if actions.RegisterAccount(obj.Username, obj.Password) {
		return c.JSON(registerResponse{
			true,
			"Successfully created account",
		})
	} else {
		return c.JSON(registerResponse{
			false,
			"failed to create an account",
		})
	}
}

func checkRegisterRequest(obj registerRequest) bool {
	return obj.Username != "" && obj.Password != ""
}
