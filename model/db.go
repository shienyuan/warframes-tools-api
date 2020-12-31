package model

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type DbError struct {
	Code    int
	Message string
}

const dbName = "warframes-tools"

var WarframeNormalCol *mongo.Collection
var WarframePrimeCol *mongo.Collection

func SetupDb() {
	client, err := mongo.NewClient(options.Client().ApplyURI(
		os.Getenv("dbAddress")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	WarframeNormalCol = client.Database(dbName).Collection("warframe-normal")
	WarframePrimeCol = client.Database(dbName).Collection("warframe-prime")

}
