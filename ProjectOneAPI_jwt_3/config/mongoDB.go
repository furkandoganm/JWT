package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ConnectionDB(dbName, cName string, isNew bool) (*mongo.Collection, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvReader("MONGOURI")))
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	if isNew {
		if _, err := client.Database(dbName).Collection(cName).InsertOne(context.Background(), nil); err != nil {
			return nil, err
		}

		return client.Database(dbName).Collection(cName), nil
	}

	return client.Database(dbName).Collection(cName), nil
}
