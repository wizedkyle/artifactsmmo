package database

import (
	"context"
	"github.com/wizedkyle/artifactsmmo/v2/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"os"
)

var (
	Client = database{}
)

type database struct {
	Client          *mongo.Client
	ItemsCollection *mongo.Collection
	TaskCollection  *mongo.Collection
}

const (
	itemsCollection = "items"
	databaseName    = "artifactsmmo"
	taskCollection  = "tasks"
)

func Init() {
	var (
		opts *options.ClientOptions
	)
	if os.Getenv("GIN_MODE") == "" {
		opts = options.Client().ApplyURI("mongodb://mongodb:27017")
	} else {
		// TODO: implement cosmosdb
	}
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		utils.Logger.Fatal("failed to connect to database", zap.Error(err))
	}
	Client.ItemsCollection = client.Database(databaseName).Collection(itemsCollection)
	Client.TaskCollection = client.Database(databaseName).Collection(taskCollection)
	Client.Client = client
}
