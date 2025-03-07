package main

import (
	"bmsp-backend-service/db"
	"bmsp-backend-service/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db.InitDB()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "*",
	}))

	h := handlers.NewHandlers()

	app.Get("/", h.RootHandler)

	app.Post("/alarm", h.CreateAlarmHandler)

	// go cronjob.StartCronjobBuzzebeeJson()
	// go cronjob.StartCronjobBuzzebeeJsonToXML()

	log.Fatal(app.Listen(":8070"))
}
