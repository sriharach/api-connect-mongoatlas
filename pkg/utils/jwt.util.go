package utils

import (
	"api-connect-mongodb-atlas/pkg/models"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type PayloadsClaims struct {
	jwt.StandardClaims
	Sub *models.ModuleProfile `json:"sub"`
}

type GenerateJWTOption struct {
	jwt.StandardClaims
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_token"`
}

func GenerateTokenJWT(payload *models.ModuleProfile, isExpired bool) (string, error) {
	var expired30m int64
	now := time.Now().UTC()

	if isExpired {
		expired30m = now.Add(30 * time.Minute).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, PayloadsClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expired30m,
			IssuedAt:  now.Unix(),
		},
		Sub: payload,
	})

	decode, _ := Decode(os.Getenv("JWT_SECRET"))

	mySigningKey := []byte((decode))

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func ValidateToken(token string) (interface{}, error) {
	decode, _ := Decode(os.Getenv("JWT_SECRET"))
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return []byte(decode), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalidate token: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, fmt.Errorf("invalid token claim")
	}

	return claims["sub"], nil
}

func ModuleUser(ctx *fiber.Ctx) *models.ModuleProfile {
	user := ctx.Locals("user").(*models.ModuleProfile)
	return user
}
