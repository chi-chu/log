package mongo

import (
	"context"
	"github.com/chi-chu/log/entry"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type printer struct {
	client				*mongo.Client
	database			string
	collection			string
}


func New(dsn, database, collection string) (*printer, error) {
	opts := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}
	return &printer{client:client, database:database, collection:collection}, nil
}

func (p *printer) Print(e *entry.Entry) {

}

func (p *printer) Rotate(b bool) error {
	return nil
}

func (p *printer) Exit() {

}