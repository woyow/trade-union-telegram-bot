package mongodb

import (
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	tradeUnionDatabase = "tradeUnion"

	usersCollection = "users"
)

type RepoImpl struct {
	db  *mongo.Client
	log *logrus.Logger
}

func NewRepoImpl(db *mongo.Client, log *logrus.Logger) *RepoImpl {
	return &RepoImpl{
		db:  db,
		log: log.WithField("mongo", "mmmm").Logger,
	}
}
