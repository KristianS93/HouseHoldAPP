package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongClient struct {
	Connection *mongo.Collection
}

func (m *MongClient) DbConnect() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(ConstDbURI))
	if err != nil {
		panic(err)
	}

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	m.Connection = client.Database(ConstDatabase).Collection(ConstCollection)
}
