package mongodb

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"trade-union-service/internal/domains/users/domain/entity"
	"trade-union-service/internal/domains/users/errs"
)

func (r *RepoImpl) GetUser(ctx context.Context, dto entity.GetUserRepoDTO) (*entity.GetUserOut, error) {
	filter := make(bson.D, 0, 1)

	if dto.ID != nil {
		objectID, err := primitive.ObjectIDFromHex(*dto.ID)
		if err != nil {
			r.log.Error("mongo: GetUser - primitive.ObjectIDFromHex error: ", err.Error())
			return nil, err
		}
		filter = append(filter, bson.E{
			Key: "_id",
			Value: objectID,
		})
	} else if dto.ChatID != nil {
		filter = append(filter, bson.E{
			Key: "chatId",
			Value: *dto.ChatID,
		})
	}

	res := r.db.Database(tradeUnionDatabase).
		Collection(usersCollection).
		FindOne(ctx, filter)

	var out entity.GetUserOut

	if err := res.Decode(&out); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.ErrUserNotFound
		}
		r.log.Error("repo: GetUser - res.Decode error: ", err.Error())
		return nil, err
	}

	return &out, nil
}
