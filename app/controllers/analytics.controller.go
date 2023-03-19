package controllers

import (
	"api-connect-mongodb-atlas/pkg/models"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OBiD = struct {
	Oid string `json:"$oid"`
}

type Items struct {
	ID         string   `json:"_id"`
	Account_id int      `json:"account_id"`
	Limit      int      `json:"limit"`
	Products   []string `json:"products"`
}

type IAnalyticsController interface {
	GetAccount(c *fiber.Ctx) error
}

type AnalyticsOptions struct {
	MongoDB *mongo.Database
}

func NewAnalyticsController(DB *mongo.Database) IAnalyticsController {
	return &AnalyticsOptions{
		MongoDB: DB,
	}
}

func (an *AnalyticsOptions) GetAccount(c *fiber.Ctx) error {
	collection := an.MongoDB.Collection("transactions")

	find, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.NewBaseErrorResponse(err.Error(), fiber.StatusInternalServerError))
	}

	var result bson.M

	for find.Next(context.Background()) {

		err := find.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
	}
	find.Close(context.Background())

	return c.JSON(result)
}
