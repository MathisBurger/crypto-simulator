package controller

import (
	"encoding/json"
	"github.com/MathisBurger/crypto-simulator/database"
	"github.com/gofiber/fiber/v2"
)

type registerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type registerResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func RegisterController(c *fiber.Ctx) error {
	raw := string(c.Body())
	obj := registerRequest{}
	err := json.Unmarshal([]byte(raw), &obj)
	if err != nil {
		return c.JSON(registerResponse{
			"failed",
			"wrong json body",
		})
	}
	if !checkRegisterRequest(obj) {
		return c.JSON(registerResponse{
			"failed",
			"wrong json body",
		})
	}
	if database.CreateAccount(obj.Username, obj.Password) {
		return c.JSON(registerResponse{
			"ok",
			"Successfully created account",
		})
	} else {
		return c.JSON(registerResponse{
			"failed",
			"failed to create an account",
		})
	}
}

func checkRegisterRequest(obj registerRequest) bool {
	return obj.Username != "" && obj.Password != ""
}
