package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/leeroyakbar/restapi-fiber/config"
	"github.com/leeroyakbar/restapi-fiber/router"
	"log"
	"os"
)

func main() {
	godotenv.Load()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	config.ConnectDB()
	router.Router(app)
	log.Fatalln(app.Listen(os.Getenv("SERVER")))
}
