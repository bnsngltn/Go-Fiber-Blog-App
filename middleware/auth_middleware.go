package middleware

import (
	"github.com/bnsngltn/go-fiber-blog-app/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

//Protected :
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:    []byte(config.Secret),
		SigningMethod: "HS512",
		ErrorHandler:  jwtError})
}

func jwtError(c *fiber.Ctx, err error) error {
	resp := fiber.Map{"message": "Missing or malformed JWT."}

	return c.Status(400).JSON(resp)
}
