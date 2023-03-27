package internal

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type HttpCookieInterface interface {
	CookieFromFiber(c *fiber.Ctx, e *DefaultParamCookie) error
}

type DefaultParamCookie struct {
	Name    string
	Value   string
	Expires int    `default:"30"`
	Path    string `default:"/"`
}

func CookieFromFiber(c *fiber.Ctx, e *DefaultParamCookie) error {
	time := time.Now().Add(time.Duration(e.Expires) * time.Minute)
	c.Cookie(&fiber.Cookie{
		Name:    e.Name,
		Value:   e.Value,
		Expires: time,
		Path:    e.Path,
	})

	return c.Next()
}
