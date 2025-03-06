package main

import (
	"bmsp-backend-service/db"
	"bmsp-backend-service/handlers"
	"bmsp-backend-service/middlewares"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",                   // เปิดให้ทุก origin สามารถเข้าถึง API ได้
		AllowMethods: "GET,POST,PUT,DELETE", // กำหนดให้รองรับ HTTP Methods ที่ต้องการ
		AllowHeaders: "*",                   // กำหนดให้รองรับ Headers
	}))

	_db := db.InitDB()

	db.Migrate(_db)

	h := handlers.NewHandlers()

	app.Get("/", handlers.RootHandler)
	app.Get("/horizon", middlewares.AuthMiddleware, h.GetListFileFromHorizon)
	// app.Get("/buzzebee", middlewares.AuthMiddleware, h.GetListFileFromBuzzebee)
	app.Get("/buzzebee/stat", middlewares.AuthMiddleware, h.GetListFileFromBuzzebeeStat)
	app.Get("/buzzebee/list", middlewares.AuthMiddleware, h.GetDocumentListFromBuzzebee)
	app.Post("/buzzebee", middlewares.AuthMiddleware, h.CreateDocumentBuzzebee)
	app.Post("/buzzebee/multiple", middlewares.AuthMiddleware, h.CreateDocumentBuzzebeeMultiple)
	app.Post("/xml/:id", middlewares.AuthMiddleware, h.CreateXML)

	app.Post("/login", h.Login)

	app.Post("/alienvaultcase", handlers.CreateAlienvaultCaseHandler)

	// go cronjob.StartCronjobBuzzebeeJson()
	// go cronjob.StartCronjobBuzzebeeJsonToXML()

	log.Fatal(app.Listen(":8070"))
}
