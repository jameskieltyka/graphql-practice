package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	host     = "127.0.0.1"
	username = "root"
	password = "password"
	database = "test"
)

func ConfigDB(ctx context.Context) (*mongo.Database, error) {
	uri := fmt.Sprintf(`mongodb://%s:%s@%s/%s`,
		username,
		password,
		host,
		database,
	)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to mongodb: %v", err)
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	db := client.Database("test")
	return db, nil
}
