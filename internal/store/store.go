package store

import (
	"context"
	"sync"
	"time"

	"github.com/zloyboy/mongo/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	Events *mongo.Collection
	mx     sync.Mutex
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
	s.mx.Lock()
	defer s.mx.Unlock()

	var result bson.M
	if err := s.Events.FindOne(ctx, bson.D{{Key: "type", Value: tp}, {Key: "state", Value: 0}}).Decode(&result); err == nil {
		return nil
	}

	_, err := s.Events.InsertOne(ctx, bson.D{
		{Key: "type", Value: tp},
		{Key: "state", Value: 0},
		{Key: "started_at", Value: time.Now().Local().Format("2006-01-02 15:04:05")},
	})
	return err
}

type Finish struct {
	Finished bool
	Error    error
}

func (s *Store) Finish(tp string) Finish {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.mx.Lock()
	defer s.mx.Unlock()

	filter := bson.D{{Key: "type", Value: tp}, {Key: "state", Value: 0}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "state", Value: 1},
		{Key: "finished_at", Value: time.Now().Local().Format("2006-01-02 15:04:05")},
	}}}
	res, err := s.Events.UpdateOne(ctx, filter, update)
	finished := false
	if res != nil && err == nil {
		if res.ModifiedCount != 0 {
			finished = true
		}
	}

	return Finish{finished, err}
}
