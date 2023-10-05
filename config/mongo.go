package config

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB(uri, dbName string) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo")
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to connect to mongo. Check if uri is correct - %s", uri)
	}

	db := client.Database(dbName)
	return db, nil
}
