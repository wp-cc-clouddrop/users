package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	. "users/internal/types"
)

var client *mongo.Client
var databaseType string

// Connect to example uri: "mongodb://localhost:27017"
func Connect(uri string) error {
	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	var err error = nil
	client, err = mongo.Connect(context.TODO(), clientOptions)

	if strings.EqualFold(uri, "mongodb://localhost:27017") {
		databaseType = "TEST"
	} else {
		databaseType = "PROD"
	}
	return err
}

func Create(collection string, newUser User) error {
	_, err := client.Database(databaseType).Collection(collection).InsertOne(context.TODO(), newUser)
	return err
}

func setIndex(fieldname string) {

}

func Disconnect() error {
	return client.Disconnect(context.TODO())
}
