package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database
var Client *mongo.Client

func InitDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	uri := os.Getenv("DATABASE_URI")
	if uri == "" {
		log.Fatal("DATABASE_URI is empty")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	Database = client.Database("go-fiber-crud")
	Client = client
}
