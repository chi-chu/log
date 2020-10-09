package mongo

import (
	"context"
	"errors"
	"fmt"
	"github.com/chi-chu/log/entry"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type printer struct {
	client				*mongo.Client
	collectionClient	*mongo.Collection
	database			string
	collection			string
	rotateCollection	string
}


func New(dsn, database, collection string) (*printer, error) {
	if database == "" {
		return nil, errors.New("database can`t be nil")
	}
	opts := options.Client().ApplyURI(dsn)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, err
	}
	return &printer{client:client, database:database, collection:collection}, nil
}

func (p *printer) Print(e *entry.Entry) {
	insertOneResult, err := p.collectionClient.InsertOne(context.Background(), e.Data)
	if err != nil {
		fmt.Println("[log] mongo insert err: ", err)
	}
	fmt.Println("collection.InsertOne: ", insertOneResult.InsertedID)
}

func (p *printer) Rotate(b bool) error {
	rotateName := p.collection
	if b {
		rotateName = p.collection + "_" + time.Now().Format("200601021504")
	}
	p.rotateCollection = rotateName
	p.collectionClient = p.client.Database(p.database).Collection(p.rotateCollection)
	return nil
}

func (p *printer) Exit() {
	_ = p.client.Disconnect(context.Background())
}