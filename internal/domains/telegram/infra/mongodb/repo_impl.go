package mongodb

import (
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// Logging.
	chatIDLoggingKey   = "chat_id"
	domainLoggingKey   = "domain"
	domainLoggingValue = "telegram"
	infraLoggingKey    = "infra"
	indraLoggingValue  = "mongodb"

	// Database.
	tradeUnionDatabase = "tradeUnion"

	// Collections.
	appealsCollection        = "appeals"
	appealSubjectsCollection = "appealSubjects"
	chatStatesCollection     = "chatStates"
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
