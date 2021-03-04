package router

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/bnsngltn/go-fiber-blog-app/docs"
	"github.com/bnsngltn/go-fiber-blog-app/handler"
	"github.com/bnsngltn/go-fiber-blog-app/middleware"
	"github.com/gofiber/fiber/v2"
)

//SetupRoutes :
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	//For Swagger Docs
	api.Get("/docs/*", swagger.Handler)

	authGroup := api.Group("/auth")
	authGroup.Post("/login", handler.Login)

	userGroup := api.Group("/user")
	userGroup.Post("/", handler.CreateUser)

	postGroup := api.Group("/post")
	postGroup.Post("/", middleware.Protected(), handler.AddPost)
	postGroup.Get("/personal", middleware.Protected(), handler.GetUserPosts)
}
