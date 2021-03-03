package main

import (
	"log"

	"gihub.com/bnsngltn/go-fiber-blog-app/database"
	"gihub.com/bnsngltn/go-fiber-blog-app/router"
	"github.com/gofiber/fiber/v2"
)

// @title Go Fiber Blog App
// @version 1.0
// @description A simple blog API that I created to practice golang
// @securityDefinitions.apikey BearerToken
// @in header
// @name Authorization
// @BasePath /api
func main() {
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	log.Println("Creating new fiber app...")
	app := fiber.New()

	log.Println("Setting up routes...")
	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
