package mongodb

import (
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	chatIDLoggingKey = "chatID"

	tradeUnionDatabase = "tradeUnion"

	chatStatesCollection = "chatStates"
)

type RepoImpl struct {
	db  *mongo.Client
	log *logrus.Logger
}

func NewRepoImpl(db *mongo.Client, log *logrus.Logger) *RepoImpl {
	return &RepoImpl{
		db:  db,
		log: log,
	}
}
