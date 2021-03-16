package main

import (
	"github.com/MathisBurger/crypto-simulator/auth"
	"github.com/MathisBurger/crypto-simulator/controller"
	"github.com/MathisBurger/crypto-simulator/database/actions"
	"github.com/MathisBurger/crypto-simulator/services"
	"github.com/MathisBurger/crypto-simulator/utils"
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

	utils.GenerateKeys()

	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		ExposeHeaders:    "Authorization",
	}))

	if os.Getenv("RATE_LIMITER") == "true" {
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
	}

	// refresh token auth controller
	app.Post("/api/auth/login", auth.LoginController)
	app.Get("/api/auth/accessToken", auth.AccessTokenController)
	app.Post("/api/auth/revokeSession", auth.RevokeSessionController)
	app.Get("/api/auth/me", auth.StatusController)

	app.Get("/api", controller.DefaultController)
	app.Post("/api/register", controller.RegisterController)
	app.Get("/api/checkBalance", controller.CheckBalanceController)
	app.Get("/api/getAllCurrencys", controller.GetAllCurrencysController)
	app.Get("/api/getCurrency", controller.GetCurrencyController)
	app.Post("/api/buyCrypto", controller.BuyCryptoController)
	app.Post("/api/sellCrypto", controller.SellCryptoController)
	app.Get("/api/getAllTrades", controller.GetAllTradesController)
	app.Get("/api/getWalletsForUser", controller.GetCryptoWalletsForUser)

	if os.Getenv("DEPRECATED_ENDPOINTS") == "true" {
		app.Get("/api/getCurrencyData", controller.GetCurrencyDataController)
		app.Post("/api/login", controller.LoginController)
		app.Get("/api/checkTokenStatus", controller.GetTokenStatusController)
	}

	// Web endpoints
	app.Static("/", "./web/dist/web")
	app.Static("/register", "./web/dist/web/index.html")
	app.Static("/login", "./web/dist/web/index.html")
	app.Static("/dashboard", "./web/dist/web/index.html")
	app.Static("/currency-view/*", "./web/dist/web/index.html")

	go services.CurrencyUpdater()
	err = app.Listen(":" + os.Getenv("APPLICATION_PORT"))
	if err != nil {
		panic(err.Error())
	}

}
