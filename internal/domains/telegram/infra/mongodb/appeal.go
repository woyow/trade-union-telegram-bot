package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"trade-union-service/internal/domains/telegram/domain/entity"
	"trade-union-service/internal/domains/users/errs"
)

const (
	appealsCollection = "appeals"
)

func (r *RepoImpl) DeleteDraftAppeal(ctx context.Context, dto entity.DeleteDraftAppealRepoDTO) error {
	filter := bson.D{{"chatId", dto.ChatID}, {"isDraft", true}}

	_, err := r.db.Database(tradeUnionDatabase).
		Collection(appealsCollection).
		DeleteOne(ctx, filter)
	if err != nil {
		r.log.WithField(chatIDLoggingKey, dto.ChatID).
			Error("mongo: DeleteDraftAppeal query error: ", err.Error())
		return err
	}

	return nil
}

func (r *RepoImpl) CreateAppeal(ctx context.Context, dto entity.CreateAppealRepoDTO) (*entity.CreateAppealOut, error) {
	doc := bson.M{
		"chatId":  dto.ChatID,
		"isDraft": dto.IsDraft,
	}

	res, err := r.db.Database(tradeUnionDatabase).
		Collection(appealsCollection).
		InsertOne(ctx, doc)
	if err != nil {
		r.log.WithField(chatIDLoggingKey, dto.ChatID).
			Error("mongo: CreateAppeal query error: ", err.Error())
		return nil, err
	}

	out := entity.CreateAppealOut{
		ID: res.InsertedID.(primitive.ObjectID).Hex(),
	}

	return &out, nil
}

func (r *RepoImpl) UpdateDraftAppeal(ctx context.Context, dto entity.UpdateAppealRepoDTO) error {
	filter := bson.D{{"chatId", dto.ChatID}, {"isDraft", true}}

	update := make(bson.D, 0, 1)

	if dto.Fname != nil {
		update = append(update, bson.E{
			Key:   "$set",
			Value: bson.D{{"fname", *dto.Fname}},
		})
	}

	if dto.Lname != nil {
		update = append(update, bson.E{
			Key:   "$set",
			Value: bson.D{{"lname", *dto.Lname}},
		})
	}

	if dto.Mname != nil {
		update = append(update, bson.E{
			Key:   "$set",
			Value: bson.D{{"mname", *dto.Mname}},
		})
	}

	if dto.IsDraft != nil {
		update = append(update, bson.E{
			Key:   "$set",
			Value: bson.D{{"isDraft", *dto.IsDraft}},
		})
	}

	_, err := r.db.Database(tradeUnionDatabase).
		Collection(appealsCollection).
		UpdateOne(ctx, filter, update)
	if err != nil {
		r.log.WithField(chatIDLoggingKey, dto.ChatID).
			Error("mongo: UpdateDraftAppeal query error: ", err.Error())
		return err
	}

	return nil
}

func (r *RepoImpl) GetDraftAppeal(ctx context.Context, dto entity.GetDraftAppealRepoDTO) (*entity.GetDraftAppealRepoOut, error) {
	filter := bson.D{{"chatId", dto.ChatID}, {"isDraft", true}}

	res := r.db.Database(tradeUnionDatabase).
		Collection(appealsCollection).
		FindOne(ctx, filter)

	var out entity.GetDraftAppealRepoOut

	if err := res.Decode(&out); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.ErrUserNotFound
		}

		r.log.WithField(chatIDLoggingKey, dto.ChatID).
			Error("repo: GetDraftAppeal - res.Decode error: ", err.Error())
		return nil, err
	}

	return &out, nil
}
