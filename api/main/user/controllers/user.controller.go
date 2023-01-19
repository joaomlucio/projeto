package controllers

import (
	""
)

func GetAll(c *fiber.Ctx) error {
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
}