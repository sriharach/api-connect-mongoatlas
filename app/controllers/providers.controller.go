package controllers

import "github.com/gofiber/fiber/v2"

type IProviders interface {
	Login(c *fiber.Ctx) error
}

func Login(c *fiber.Ctx) error {
	return c.SendString("ok")
}
