package controllers

import (
	"api-connect-mongodb-atlas/pkg/models"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IuserController interface {
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

func (ur *PropsUserController) GetUserAccount(c *fiber.Ctx) error {
	user_id := c.Cookies("user_id")

	if user_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.NewBaseErrorResponse("user_id not found.", fiber.StatusBadRequest))
	}
	var result models.ModuleProfile

	collection := ur.MainCollectionDB

	docID, _ := primitive.ObjectIDFromHex(user_id)
	bson := bson.M{"_id": docID}

	err := collection.FindOne(context.Background(), bson).Decode(&result)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.NewBaseErrorResponse(err.Error(), fiber.StatusBadRequest))
	}
	return c.JSON(models.NewBaseResponse(result, fiber.StatusOK))
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
