package store

import (
	"context"
	"time"

	"github.com/zloyboy/mongo/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	Events *mongo.Collection
}

func New(config *config.Config) (*Store, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoURL))
	if err == nil && client != nil {
		err = client.Ping(ctx, nil)
	}
	if err != nil {
		return nil, err
	}

	return &Store{Events: client.Database(config.MongoDB).Collection("events")}, nil
}

func (s *Store) Start(tp string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := s.Events.InsertOne(ctx, bson.D{{Key: "type", Value: tp}, {Key: "state", Value: 0}})
	return err
}

func (s *Store) Finish() error {
	return nil
}
