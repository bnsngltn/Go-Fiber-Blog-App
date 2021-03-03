package router

import (
	_ "gihub.com/bnsngltn/go-fiber-blog-app/docs"
	"gihub.com/bnsngltn/go-fiber-blog-app/handler"
	"gihub.com/bnsngltn/go-fiber-blog-app/middleware"
	swagger "github.com/arsmn/fiber-swagger/v2"
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
