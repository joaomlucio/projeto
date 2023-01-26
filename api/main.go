package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/joaomlucio/projeto/api/user/controllers"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/api/v1/user", controllers.GetAll)

	app.Get("/api/v1/user/:id", controllers.GetOne)

	app.Post("/api/v1/user", controllers.CreateUser)

	app.Put("/api/v1/user/:id", controllers.UpdateUser)

	app.Delete("/api/v1/user/:id", controllers.DeleteUser)

	app.Listen(":3000")
}
