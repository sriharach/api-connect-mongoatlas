package middleware

import (
	"api-connect-mongodb-atlas/pkg/models"
	"api-connect-mongodb-atlas/pkg/utils"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func DeserializeUser(c *fiber.Ctx) error {
	str, is_jwt := PassportJwtValidate(c)
	if is_jwt {
		return c.Status(fiber.StatusUnauthorized).JSON(models.NewBaseErrorResponse(fiber.Map{
			"message": str,
		}, fiber.StatusUnauthorized))
	}

	return c.Next()
}

func PassportJwtValidate(c *fiber.Ctx) (string, bool) {
	decode, _ := utils.Decode(os.Getenv("JWT_SECRET"))

	// bearer := c.Get("Authorization")
	bearer := c.Cookies("access_token")

	if bearer == "" {
		return "Unauthorized", true
	}

	// trimToken := strings.TrimPrefix(bearer, "Bearer ")

	_, err := jwt.ParseWithClaims(bearer, &utils.PayloadsClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(decode), nil
	})

	if err != nil {
		// c.Locals("user", nil)
		return err.Error(), true
	}
	return "is_error", false
}
