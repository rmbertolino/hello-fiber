package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
)


type User struct{
	Id string
	Username string
	Name string
}

func handleUsers(c *fiber.Ctx) error {
	user := User{
		Username: "U190804",
		Name: "Rodolfo Bertolino",
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func handleCreateUser(c *fiber.Ctx) error {
	user := User{}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.Id = uuid.NewString()

	return c.Status(fiber.StatusCreated).JSON(user)
}

func main(){

	app := fiber.New()

	//Middlewares
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		data := map[string]interface{}{
			"message": "Hello, World!!!!",
			"status":  "success",
		}

		return c.JSON(data)
	})

	app.Use(requestid.New()) //añade id en el header del response todas las peticiones que vienen debajo
	userGroup := app.Group("/users")
	userGroup.Get("", handleUsers)
	userGroup.Post("", handleCreateUser)


	app.Listen(":3000")
}