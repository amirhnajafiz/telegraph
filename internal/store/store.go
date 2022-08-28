package store

import "go.mongodb.org/mongo-driver/mongo"

type Store struct {
	Database *mongo.Database
}
