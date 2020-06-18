package db

import (
	"context"
	"time"

	"github.com/beldmian/ORS/internal/app/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Datebase ...
type Datebase struct {
	URI string
}

// New ...
func New(uri string) Datebase {
	var db Datebase
	db.URI = uri

	return db
}

// GetEvents ...
func (db Datebase) GetEvents() ([]types.Event, error) {
	ctx5, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	Client, err := mongo.Connect(ctx5, options.Client().ApplyURI(db.URI))
	if err != nil {
		return nil, err
	}
	eventsCollection := Client.Database("ors").Collection("events")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := eventsCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var events []types.Event
	for cur.Next(ctx) {
		var result types.Event
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		events = append(events, result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return events, nil
}
