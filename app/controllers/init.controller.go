package controllers

import "github.com/gofiber/fiber/v2"

func TestConnect(c *fiber.Ctx) error {
	return c.SendString("Welcome to Golang.")
}
