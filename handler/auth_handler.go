package handler

import (
	"gihub.com/bnsngltn/go-fiber-blog-app/config"
	"gihub.com/bnsngltn/go-fiber-blog-app/database"
	"gihub.com/bnsngltn/go-fiber-blog-app/model"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type loginStruct struct {
	Username string
	Password string
}

//Login handler
func Login(c *fiber.Ctx) error {
	login := new(loginStruct)
	user := new(model.User)
	var err error

	if err = c.BodyParser(login); err != nil {
		resp := fiber.Map{"message": "Bad request."}
		return c.Status(400).JSON(resp)
	}

	if user, err = GetUserByUsername(login.Username); err != nil {
		resp := fiber.Map{"message": "Username not found."}
		return c.Status(404).JSON(resp)
	}

	if err = verifyPassword(login.Password, user.Password); err != nil {
		resp := fiber.Map{"message": "Incorrect password."}
		return c.Status(401).JSON(resp)
	}

	var token string
	token, err = generateJWTToken(user)
	if err != nil {
		resp := fiber.Map{"message": "Something went wrong."}
		return c.Status(500).JSON(resp)
	}

	resp := fiber.Map{"message": "Successfully logged in!", "token": token}
	return c.JSON(resp)
}

//GetUserByUsername func
func GetUserByUsername(username string) (*model.User, error) {
	db := database.DB

	var user model.User
	var err error

	if err = db.Where(&model.User{Username: username}).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil

}

func verifyPassword(password, passwordHash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))

	return err
}

func generateJWTToken(user *model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)

	// This is the only line of code here in which I cannot understand
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username

	t, err := token.SignedString([]byte(config.Secret))

	return t, err
}
