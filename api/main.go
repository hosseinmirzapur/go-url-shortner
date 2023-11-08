package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hosseinmirzapur/urlshortner/routes"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get(":url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	app := fiber.New()

	app.Use(logger.New())
	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
