package routes

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserModel struct {
	Id       int32
	Username string
	Email    string
}

var users = []UserModel{
	{Id: 0, Username: "Alex", Email: "Smith@Smith.com"},
}

func createUser(c *fiber.Ctx) error {
	var user UserModel
	c.BodyParser(&user)
	users = append(users, user)

	return c.SendString(fmt.Sprintf("%v", users))
}

func getUser(c *fiber.Ctx) error {
	id, ok := c.Queries()["id"]
	index, err := strconv.Atoi(id)

	if !ok || (err != nil) || index < 0 || index > len(users) {
		return fiber.NewError(fiber.StatusBadRequest, "Unknown request.")
	}
	return c.JSON(users[index])
}

func User() *fiber.App {
	userRoute := fiber.New()
	userRoute.Post("/", createUser)
	userRoute.Get("/", getUser)
	return userRoute
}
