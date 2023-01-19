package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"context"
	"encoding/json"
)

// func init() {
// 	//clientOptions := options.Client().ApplyURI("mongodb://user:password@mongo:27017/")
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
// 	client, err := mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = client.Ping(ctx, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	collection = client.Database("local").Collection("projeto")
// }

func main() {
	app := fiber.New()

	app.Get("/api/user", func(c *fiber.Ctx) error {
		users, err := findUsers()
		if err != nil {
			return c.Status(404).SendString(err.Error())
		}
		
		blob, err := json.Marshal(users)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		
		c.Response().BodyWriter().Write(blob)
		return c.SendStatus(200)
	})

	app.Post("/api/user", func(c *fiber.Ctx) error {
		user := new(CreateUser)
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			}) 
		}

		errors := ValidateStruct(*user)
		if errors != nil {
       		return c.Status(fiber.StatusBadRequest).JSON(errors)        
    	}

		u, err := createUser(user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			}) 
		}
		return c.Status(201).JSON(u)
	})

	app.Patch("/api/user/:id", func(c *fiber.Ctx) error {
		user := new(UpdateUser)
		id := c.Params("id")
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			}) 
		}

		errors := ValidateStruct(*user)
		if errors != nil {
       		return c.Status(fiber.StatusBadRequest).JSON(errors)        
    	}

		u, err := updateUser(id, user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			}) 
		}
		return c.Status(200).JSON(u)
	})

	app.Delete("/api/user/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		u, err := deleteUser(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			}) 
		}
		return c.Status(200).JSON(u)
	})

	app.Listen(":3000")
}
