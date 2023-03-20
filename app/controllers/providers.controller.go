package controllers

import (
	"api-connect-mongodb-atlas/pkg/models"
	"api-connect-mongodb-atlas/pkg/utils"
	"context"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IProviders interface {
	Login(c *fiber.Ctx) error
}

type PropsProviderController struct {
	MongoDB          *mongo.Database
	MainCollectionDB *mongo.Collection
}

func NewProviderControllers(DB *mongo.Database) IProviders {
	return &PropsProviderController{
		MongoDB:          DB,
		MainCollectionDB: DB.Collection("users"),
	}
}

func (pv *PropsProviderController) Login(c *fiber.Ctx) error {
	var result *models.ModuleProfile
	payload := new(models.SignInInput)
	collection := pv.MainCollectionDB

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.NewBaseErrorResponse(err.Error(), fiber.StatusBadRequest))
	}
	bson := bson.M{
		"e_mail": payload.E_mail,
	}
	err := collection.FindOne(context.Background(), bson).Decode(&result)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.NewBaseErrorResponse(err.Error(), fiber.StatusNotFound))
	}

	if is_passwor_hash := utils.CheckPasswordHash(payload.Password, result.Password); !is_passwor_hash {
		return c.Status(fiber.StatusNotAcceptable).JSON(models.NewBaseErrorResponse("Password don't matching", fiber.StatusNotAcceptable))
	}

	access_token, _ := utils.GenerateTokenJWT(result, true)
	refresh_token, _ := utils.GenerateTokenJWT(result, false)

	decode, _ := utils.Decode(os.Getenv("JWT_SECRET"))
	token, _err := jwt.ParseWithClaims(access_token, &utils.PayloadsClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(decode), nil
	})

	if _err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.NewBaseErrorResponse(_err.Error(), fiber.StatusBadRequest))
	}

	claims := token.Claims.(*utils.PayloadsClaims)

	return c.JSON(models.NewBaseResponse(utils.GenerateJWTOption{
		Access_token:  access_token,
		Refresh_token: refresh_token,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: claims.ExpiresAt,
			IssuedAt:  claims.IssuedAt,
		},
	}, fiber.StatusOK))
}
