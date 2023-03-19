package middleware

import (
	"api-connect-mongodb-atlas/pkg/models"
	"api-connect-mongodb-atlas/pkg/utils"
	"os"
	"strings"

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

	bearer := c.Get("Authorization")
	if bearer == "" {
		c.Locals("user", nil)
		return "Unauthorized", true
	}

	trimToken := strings.TrimPrefix(bearer, "Bearer ")

	// if _, err := utils.ValidateToken(trimToken); err != nil {
	// 	c.Locals("user", nil)
	// 	return "validate", true
	// }

	token, err := jwt.ParseWithClaims(trimToken, &utils.PayloadsClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(decode), nil
	})

	if err != nil {
		c.Locals("user", nil)
		return err.Error(), true
	}

	claims := token.Claims.(*utils.PayloadsClaims)

	c.Locals("user", claims.Sub)

	return "is_error", false

}
