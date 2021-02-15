package main

import (
	"github.com/MathisBurger/crypto-simulator/controller"
	"github.com/MathisBurger/crypto-simulator/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	database.CreateRequiredTables()

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Get("/", controller.DefaultController)

	err = app.Listen(":" + os.Getenv("APPLICATION_PORT"))
	if err != nil {
		panic(err.Error())
	}

}
