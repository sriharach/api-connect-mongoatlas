package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InintMongodbAtlas() *mongo.Database {
	// Set up client options

	var (
		MONGODB_URI    = os.Getenv("MONGODB_URI")
		MONGODB_DBNAME = os.Getenv("MONGODB_DBNAME")
	)

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(MONGODB_URI).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// var database = client.Database(os.Getenv("MONGODB_DBNAME"))
	var database = client.Database(MONGODB_DBNAME)
	fmt.Println("Connect to Mongo Database Atlas.")

	return database
}
