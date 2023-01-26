package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/joaomlucio/projeto/api/user/models/dtos"
	"github.com/joaomlucio/projeto/api/user/services"
)

func GetAll(c *fiber.Ctx) error {
	users, err := services.FindUsers()
	if err != nil {
		var e *fiber.Error
        if errors.As(err, &e) {
            return c.Status(e.Code).SendString(e.Message)
        }
		return c.Status(500).SendString(err.Error())
	}
	
	return c.Status(200).JSON(users)
}

func GetOne(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := services.FindUser(id)
	if err != nil {
		var e *fiber.Error
        if errors.As(err, &e) {
            return c.Status(e.Code).SendString(e.Message)
        }
		return c.Status(500).SendString(err.Error())
	}
	
	return c.Status(200).JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(dtos.CreateUser)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		}) 
	}

	errs := services.ValidateStruct(*user)
	if errs != nil {
		   return c.Status(fiber.StatusBadRequest).JSON(errs)        
	}

	u, err := services.CreateUser(user)
	if err != nil {
		var e *fiber.Error
        if errors.As(err, &e) {
            return c.Status(e.Code).SendString(e.Message)
        }
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(201).JSON(u)
}

func UpdateUser(c *fiber.Ctx) error {
	user := new(dtos.UpdateUser)
	id := c.Params("id")
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		}) 
	}

	errs := services.ValidateStruct(*user)
	if errs != nil {
		   return c.Status(fiber.StatusBadRequest).JSON(errs)        
	}

	u, err := services.UpdateUser(id, user)
	if err != nil {
		var e *fiber.Error
        if errors.As(err, &e) {
            return c.Status(e.Code).SendString(e.Message)
        }
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).JSON(u)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	err := services.DeleteUser(id)
	if err != nil {
		var e *fiber.Error
        if errors.As(err, &e) {
            return c.Status(e.Code).SendString(e.Message)
        }
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(204).SendString("User Deleted")
}


