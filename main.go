package main

import (
	"github.com/MathisBurger/crypto-simulator/controller"
	"github.com/MathisBurger/crypto-simulator/database/actions"
	"github.com/MathisBurger/crypto-simulator/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	actions.InitTables()

	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        20,
		Expiration: 10 * time.Second,
		KeyGenerator: func(ctx *fiber.Ctx) string {
			return ctx.Get("x-forwarded-for")
		},
		LimitReached: func(ctx *fiber.Ctx) error {
			return ctx.SendStatus(fiber.StatusTooManyRequests)
		},
	}))

	app.Get("/api", controller.DefaultController)
	app.Post("/api/register", controller.RegisterController)
	app.Post("/api/login", controller.LoginController)
	app.Get("/api/checkBalance", controller.CheckBalanceController)
	app.Get("/api/getAllCurrencys", controller.GetAllCurrencysController)
	app.Get("/api/getCurrencyData", controller.GetCurrencyDataController)
	app.Get("/api/getCurrency", controller.GetCurrencyController)

	go services.CurrencyUpdater()
	err = app.Listen(":" + os.Getenv("APPLICATION_PORT"))
	if err != nil {
		panic(err.Error())
	}

}
