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

//DbConnect has to be called after instantiating a new connection to establish the connection to mongo
func (m *MongClient) DbConnect(Collection string) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(ConstDbURI))
	if err != nil {
		panic(err)
	}

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	m.Connection = client.Database(ConstDatabase).Collection(Collection)
}
