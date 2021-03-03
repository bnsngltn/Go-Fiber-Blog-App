package handler

import (
	"gihub.com/bnsngltn/go-fiber-blog-app/database"
	"gihub.com/bnsngltn/go-fiber-blog-app/model"
	"github.com/gofiber/fiber/v2"
)

// AddPost godoc
// @Tags post
// @Summary Add a post
// @Description Add a post
// @Accept json
// @Produce json
// @Router /post [post]
func AddPost(c *fiber.Ctx) error {
	user, err := GetJWTIdentity(c)
	if err != nil {
		resp := fiber.Map{"message": "Cannot find associated user with your account."}
		return c.Status(404).JSON(resp)
	}

	post := new(model.Post)
	if err := c.BodyParser(post); err != nil {
		resp := fiber.Map{"message": "Bad request"}
		return c.Status(400).JSON(resp)
	}

	post.UserID = int(user.ID)

	db := database.DB
	if res := db.Create(&post); res.Error != nil {
		resp := fiber.Map{"message": "Title must be unique."}
		return c.Status(400).JSON(resp)
	}

	return c.JSON(post)
}

// GetUserPosts godoc
// @Tags post
// @Summary Gets all posts of a user
// @Description Gets all posts made by the user
// @Accept json
// @Produce json
// @Security BearerToken
// @Success 200 {array} model.Post
// @Router /post/personal [get]
func GetUserPosts(c *fiber.Ctx) error {
	user, err := GetJWTIdentity(c)
	if err != nil {
		resp := fiber.Map{"message": "Cannot find associated user with your account."}
		return c.Status(404).JSON(resp)
	}

	var posts []model.Post
	db := database.DB
	if err := db.Where(&model.Post{UserID: int(user.ID)}).Find(&posts).Error; err != nil {
		resp := fiber.Map{"message": "Something went wrong"}
		return c.Status(500).JSON(resp)
	}

	return c.JSON(posts)
}
