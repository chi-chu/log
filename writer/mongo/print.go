package mongo

import (
	"context"
	"fmt"
	"github.com/chi-chu/log/entry"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

type printer struct {
	client				*mongo.Client
	databaseClient		*mongo.Database
	database			string
	collection			string
	rotateCollection	string
}


func New(dsn, database, collection string) (*printer, error) {
	opts := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, err
	}
	fmt.Println(client.Database(database).Name())
	return &printer{client:client, databaseClient:client.Database(database),
		database:database, collection:collection}, nil
}

func (p *printer) Print(e *entry.Entry) {
	insertOneResult, err := p.databaseClient.Collection(p.rotateCollection).InsertOne(context.Background(), e.Data)
	if err != nil {
		fmt.Println("[log] mongo insert err: ", err)
	}
	log.Println("collection.InsertOne: ", insertOneResult.InsertedID)
}

func (p *printer) Rotate(b bool) error {
	return nil
}

func (p *printer) Exit() {
	_ = p.client.Disconnect(context.Background())
}