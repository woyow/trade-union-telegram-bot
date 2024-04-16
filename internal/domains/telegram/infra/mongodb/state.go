package mongodb

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"trade-union-service/internal/domains/telegram/domain/entity"
	"trade-union-service/internal/domains/telegram/errs"
)

const (
	errChatStateAlreadyExists = "tradeUnion.chatStates index: chatId_1 dup key"
)

func (r *RepoImpl) CreateChatCurrentState(ctx context.Context, dto entity.CreateChatCurrentStateRepoDTO) (*entity.CreateChatCurrentStateOut, error) {
	doc := bson.M{
		"chatId": dto.ChatID,
		"state":  dto.State,
	}

	res, err := r.db.Database(tradeUnionDatabase).
		Collection(chatStatesCollection).
		InsertOne(ctx, doc)
	if err != nil {
		if strings.Contains(err.Error(), errChatStateAlreadyExists) {
			return nil, errs.ErrChatCurrentStateAlreadyExists
		}

		r.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			domainLoggingKey: domainLoggingValue,
			infraLoggingKey:  indraLoggingValue,
		}).Error("CreateChatCurrentState query error: ", err.Error())

		return nil, err
	}

	out := entity.CreateChatCurrentStateOut{
		ID: res.InsertedID.(primitive.ObjectID).Hex(),
	}

	return &out, nil
}

func (r *RepoImpl) SetChatCurrentState(ctx context.Context, dto entity.SetChatCurrentStateRepoDTO) error {
	filter := bson.D{{"chatId", dto.ChatID}}
	update := bson.D{{"$set", bson.D{{"state", dto.State}}}}

	_, err := r.db.Database(tradeUnionDatabase).
		Collection(chatStatesCollection).
		UpdateOne(ctx, filter, update)
	if err != nil {
		r.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			domainLoggingKey: domainLoggingValue,
			infraLoggingKey:  indraLoggingValue,
		}).Error("SetChatCurrentState query error: ", err.Error())

		return err
	}

	return nil
}

func (r *RepoImpl) GetChatCurrentState(ctx context.Context, dto entity.GetChatCurrentStateRepoDTO) (*entity.GetChatCurrentStateOut, error) {
	filter := bson.D{{"chatId", dto.ChatID}}

	res := r.db.Database(tradeUnionDatabase).
		Collection(chatStatesCollection).
		FindOne(ctx, filter)

	var out entity.GetChatCurrentStateOut

	if err := res.Decode(&out); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.ErrChatCurrentStateNotExists
		}

		r.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			domainLoggingKey: domainLoggingValue,
			infraLoggingKey:  indraLoggingValue,
		}).Error("GetChatCurrentState - res.Decode error: ", err.Error())

		return nil, err
	}

	return &out, nil
}
