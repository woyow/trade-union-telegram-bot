package mongodb

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"trade-union-service/internal/domains/telegram/domain/entity"
	"trade-union-service/internal/domains/users/errs"
)

func (r *RepoImpl) DeleteDraftAppeal(ctx context.Context, dto entity.DeleteDraftAppealRepoDTO) error {
	filter := bson.D{{"chatId", dto.ChatID}, {"isDraft", true}}

	_, err := r.db.Database(tradeUnionDatabase).
		Collection(appealsCollection).
		DeleteOne(ctx, filter)
	if err != nil {
		r.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			domainLoggingKey: domainLoggingValue,
			infraLoggingKey:  indraLoggingValue,
		}).Error("DeleteDraftAppeal query error: ", err.Error())

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
		r.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			domainLoggingKey: domainLoggingValue,
			infraLoggingKey:  indraLoggingValue,
		}).Error("CreateAppeal query error: ", err.Error())

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

	if dto.FullName != nil {
		update = append(update, bson.E{
			Key:   "$set",
			Value: bson.D{{"fullName", *dto.FullName}},
		})
	}

	if dto.Subject != nil {
		update = append(update, bson.E{
			Key:   "$set",
			Value: bson.D{{"subject", *dto.Subject}},
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
		r.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			domainLoggingKey: domainLoggingValue,
			infraLoggingKey:  indraLoggingValue,
		}).Error("UpdateDraftAppeal query error: ", err.Error())

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

		r.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			domainLoggingKey: domainLoggingValue,
			infraLoggingKey:  indraLoggingValue,
		}).Error("GetDraftAppeal - res.Decode error: ", err.Error())

		return nil, err
	}

	return &out, nil
}

func (r *RepoImpl) GetAppealSubjects(ctx context.Context, dto entity.GetAppealSubjectsRepoDTO) (entity.GetAppealSubjectsRepoOut, error) {
	filter := make(bson.D, 0, 1)

	if dto.IsActive != nil {
		filter = append(filter, bson.E{
			Key:   "isActive",
			Value: *dto.IsActive,
		})
	}

	cur, err := r.db.Database(tradeUnionDatabase).
		Collection(appealSubjectsCollection).
		Find(ctx, filter)
	if err != nil {
		r.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			domainLoggingKey: domainLoggingValue,
			infraLoggingKey:  indraLoggingValue,
		}).Error("GetAppealSubjects query error: ", err.Error())

		return nil, err
	}

	var out entity.GetAppealSubjectsRepoOut

	if err := cur.All(ctx, &out); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}

		r.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			domainLoggingKey: domainLoggingValue,
			infraLoggingKey:  indraLoggingValue,
		}).Error("GetAppealSubjects - res.Decode error: ", err.Error())

		return nil, err
	}

	return out, nil
}
