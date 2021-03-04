package handler

import (
	"github.com/bnsngltn/go-fiber-blog-app/model"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

//GetJWTIdentity :
func GetJWTIdentity(c *fiber.Ctx) (*model.User, error) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	u, err := GetUserByUsername(username)

	return u, err
}
