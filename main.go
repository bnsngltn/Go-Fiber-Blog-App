package main

import (
	"log"

	"gihub.com/bnsngltn/go-fiber-blog-app/database"
	"gihub.com/bnsngltn/go-fiber-blog-app/router"
	"github.com/gofiber/fiber/v2"
)

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
