package mongodb

import (
	"context"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"trade-union-service/internal/domains/users/domain/entity"
	"trade-union-service/internal/domains/users/errs"
)

const (
	errUserWithChatIDAlreadyExists = "tradeUnion.users index: chatId_1 dup key"
)

func (r *RepoImpl) CreateUser(ctx context.Context, dto entity.CreateUserRepoDTO) (*entity.CreateUserOut, error) {
	doc := bson.M{
		"chatId":   dto.ChatID,
		"fname":    dto.Fname,
		"lname":    dto.Lname,
		"mname":    dto.Mname,
		"position": dto.Position,
		"roles":    dto.Roles,
	}

	res, err := r.db.Database(tradeUnionDatabase).
		Collection(usersCollection).
		InsertOne(ctx, doc)
	if err != nil {
		if strings.Contains(err.Error(), errUserWithChatIDAlreadyExists) {
			return nil, errs.ErrUserWithChatIDAlreadyExists
		}
		r.log.Error("mongo: CreateUser query error: ", err.Error())
		return nil, err
	}

	out := entity.CreateUserOut{
		ID: res.InsertedID.(primitive.ObjectID).Hex(),
	}

	return &out, nil
}
