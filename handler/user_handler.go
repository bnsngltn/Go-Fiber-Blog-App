package handler

import (
	"gihub.com/bnsngltn/go-fiber-blog-app/database"
	"gihub.com/bnsngltn/go-fiber-blog-app/model"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

//CreateUser func
func CreateUser(c *fiber.Ctx) error {
	user := new(model.User)
	var err error

	if err = c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	//Hash password
	user.Password, err = hashPassword(user.Password)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	db := database.DB
	if err := db.Create(&user).Error; err != nil {
		resp := fiber.Map{"message": "Username must be unique."}
		return c.Status(400).JSON(resp)
	}

	// When a user is created I dont want to send back the password
	// I tried all kinds of commands withint the User model
	//such as `json:"-"` but this seems to be producing undesired results
	//So I'll just define an anonymous struct here lol
	//Life hack
	u := struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	}{int(user.ID), user.Username}

	return c.JSON(u)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	return string(bytes), err
}
