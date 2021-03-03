package handler

import (
	"gihub.com/bnsngltn/go-fiber-blog-app/database"
	"gihub.com/bnsngltn/go-fiber-blog-app/model"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser godoc
// @Tags user
// @Summary Create a user
// @Description Create a user account in the DB
// @Accept json
// @Produce json
// @Param payload body model.CreateUserStruct true "The payload."
// @Success 200 {object} model.UserJSON
// @Router /user [post]
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

	return c.JSON(user.JSON())
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	return string(bytes), err
}
