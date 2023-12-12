package controllers

import (
	"api-connect-mongodb-atlas/internal"
	"api-connect-mongodb-atlas/pkg/models"
	"api-connect-mongodb-atlas/pkg/utils"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IProviders interface {
	Login(c *fiber.Ctx) error
	Oauth2(c *fiber.Ctx) error
	OauthCallback(c *fiber.Ctx) error
	RegisterAccount(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
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

func (ur *PropsProviderController) RegisterAccount(c *fiber.Ctx) error {
	requestUser := new(models.ModuleProfile)
	collection := ur.MainCollectionDB

	// parse the request body and bind it to the user instance
	if err := c.BodyParser(&requestUser); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.NewBaseErrorResponse(err.Error(), fiber.StatusBadRequest))
	}

	hashPassword, _ := utils.HashPassword(requestUser.Password)
	requestUser.Is_online = false
	requestUser.Password = hashPassword
	requestUser.ID = primitive.NewObjectID()

	_, err := collection.InsertOne(context.Background(), requestUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.NewBaseErrorResponse(err.Error(), fiber.StatusBadRequest))
	}

	return c.Status(fiber.StatusCreated).JSON(models.NewBaseResponse(requestUser, fiber.StatusCreated))
}

func (pv *PropsProviderController) Login(c *fiber.Ctx) error {
	var result *models.ModuleProfile
	payload := new(models.SignInInput)
	collection := pv.MainCollectionDB

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.NewBaseErrorResponse(err.Error(), fiber.StatusBadRequest))
	}
	e_mail := bson.M{
		"e_mail": payload.E_mail,
	}

	err := collection.FindOne(context.Background(), e_mail).Decode(&result)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.NewBaseErrorResponse(err.Error(), fiber.StatusNotFound))
	}

	if is_passwor_hash := utils.CheckPasswordHash(payload.Password, result.Password); !is_passwor_hash {
		return c.Status(fiber.StatusNotAcceptable).JSON(models.NewBaseErrorResponse("Password don't matching", fiber.StatusNotAcceptable))
	}

	id, _ := primitive.ObjectIDFromHex(c.Cookies("user_id"))
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"is_online", true}}}}

	updated, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.NewBaseErrorResponse("Matched %v documents and modified %v documents.\n", fiber.StatusBadRequest))
	}
	fmt.Printf("Matched %v documents and modified %v documents.\n", updated.MatchedCount, updated.ModifiedCount)

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

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    access_token,
		Path:     "/",
		Expires:  time.Now().Add(20 * time.Minute),
		HTTPOnly: true,
	})
	c.Cookie(&fiber.Cookie{
		Name:     "user_id",
		Value:    result.ID.Hex(),
		Path:     "/",
		Expires:  time.Now().Add(20 * time.Minute),
		HTTPOnly: true,
	})

	return c.JSON(models.NewBaseResponse(utils.GenerateJWTOption{
		Access_token:  access_token,
		Refresh_token: refresh_token,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: claims.ExpiresAt,
			IssuedAt:  claims.IssuedAt,
		},
	}, fiber.StatusOK))
}

func (pv *PropsProviderController) Oauth2(c *fiber.Ctx) error {
	confixOauth := internal.Oauth()
	url := confixOauth.AuthCodeURL("")

	return c.JSON(models.NewBaseResponse(url, fiber.StatusOK))
}

func (pv *PropsProviderController) OauthCallback(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{})
}

func (pv *PropsProviderController) Logout(c *fiber.Ctx) error {
	collection := pv.MainCollectionDB
	access_token := c.Cookies("access_token")
	fmt.Println(access_token)
	if access_token == "" {
		return c.Status(fiber.StatusNotFound).JSON(models.NewBaseErrorResponse("Not found access_token.", fiber.StatusNotFound))
	}

	id, _ := primitive.ObjectIDFromHex(c.Cookies("user_id"))
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"is_online", false}}}}

	updated, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.NewBaseErrorResponse(err.Error(), fiber.StatusBadRequest))
	}
	c.Cookie(&fiber.Cookie{
		Name:     "cookie",
		Expires:  time.Now().Add(-(time.Hour * 2)),
		HTTPOnly: true,
	})
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Path:     "/",
		Expires:  time.Now().Add(-(time.Hour * 2)),
		HTTPOnly: true,
	})
	c.Cookie(&fiber.Cookie{
		Name:     "user_id",
		Path:     "/",
		Expires:  time.Now().Add(-(time.Hour * 2)),
		HTTPOnly: true,
	})

	fmt.Printf("Matched %v documents and modified %v documents.\n", updated.MatchedCount, updated.ModifiedCount)

	return c.SendStatus(200)
}
