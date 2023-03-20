package controllers

import (
	"api-connect-mongodb-atlas/pkg/models"
	"api-connect-mongodb-atlas/pkg/utils"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IuserController interface {
	RegisterAccount(c *fiber.Ctx) error
	GetUserAccount(c *fiber.Ctx) error
	GetUsersAccount(c *fiber.Ctx) error
}

type PropsUserController struct {
	MongoDB          *mongo.Database
	MainCollectionDB *mongo.Collection
}

func NewUserControllers(DB *mongo.Database) IuserController {
	return &PropsUserController{
		MongoDB:          DB,
		MainCollectionDB: DB.Collection("users"),
	}
}

type response1 struct {
	Page   int
	Fruits []string
}

func (ur *PropsUserController) RegisterAccount(c *fiber.Ctx) error {
	requestUser := new(models.ModuleProfile)
	collection := ur.MainCollectionDB

	// parse the request body and bind it to the user instance
	if err := c.BodyParser(requestUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.NewBaseErrorResponse(err.Error(), fiber.StatusBadRequest))
	}

	hashPassword, _ := utils.HashPassword(requestUser.Password)
	requestUser.Password = hashPassword

	res, err := collection.InsertOne(context.Background(), requestUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.NewBaseErrorResponse(err.Error(), fiber.StatusBadRequest))
	}
	id := res.InsertedID

	return c.JSON(models.NewBaseResponse(id, fiber.StatusOK))
}

func (ur *PropsUserController) GetUserAccount(c *fiber.Ctx) error {
	var result models.ModuleProfile

	collection := ur.MainCollectionDB

	e_mail := c.Query("e_mail")

	bson := bson.M{
		"e_mail": e_mail,
	}

	err := collection.FindOne(context.Background(), bson).Decode(&result)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.NewBaseErrorResponse(err.Error(), fiber.StatusBadRequest))
	} else {
		return c.JSON(models.NewBaseResponse(result, fiber.StatusOK))
	}
}

func (ur *PropsUserController) GetUsersAccount(c *fiber.Ctx) error {
	collection := ur.MainCollectionDB

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	var results []models.ModuleProfile
	for cursor.Next(context.Background()) {
		var bson models.ModuleProfile
		err := cursor.Decode(&bson)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, bson)
	}

	return c.JSON(models.NewBaseResponse(results, fiber.StatusOK))
}
