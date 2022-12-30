package helpers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var client *mongo.Client

func ConnectMongodb() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MAGOBOT_MONGODB_URI")))

	if err != nil {
		panic(err)
	}
}

func DisconnectMongodb() {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func GetMongoCollection(collection string) *mongo.Collection {
	return client.Database(os.Getenv("MAGOBOT_MONGODB_DATABASE")).Collection(collection)
}
