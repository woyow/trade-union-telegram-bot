package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"

	"trade-union-service/internal/domains/users/domain/entity"
	"trade-union-service/internal/domains/users/errs"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *RepoImpl) UpdateUser(ctx context.Context, dto entity.UpdateUserRepoDTO) error {
	objectID, err := primitive.ObjectIDFromHex(dto.ID)
	if err != nil {
		r.log.Error("mongo: UpdateUser - primitive.ObjectIDFromHex error: ", err.Error())
		return err
	}

	filter := bson.D{{"_id", objectID}}

	update := make(bson.D, 0, 6)

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

	if dto.Position != nil {
		update = append(update, bson.E{
			Key:   "$set",
			Value: bson.D{{"position", *dto.Position}},
		})
	}

	if dto.ChatID != nil {
		update = append(update, bson.E{
			Key:   "$set",
			Value: bson.D{{"chatId", *dto.ChatID}},
		})
	}

	if dto.Roles != nil {
		update = append(update, bson.E{
			Key:   "$set",
			Value: bson.D{{"chatId", dto.Roles}},
		})
	}

	if len(update) == 0 {
		return errs.ErrFieldRequiredForUpdate
	}

	if _, err := r.db.Database(tradeUnionDatabase).
		Collection(usersCollection).
		UpdateOne(ctx, filter, update); err != nil {
		if strings.Contains(err.Error(), errUserWithChatIDAlreadyExists) {
			return errs.ErrUserWithChatIDAlreadyExists
		}
		r.log.Error("mongo: UpdateUser query error: ", err.Error())
		return err
	}

	return nil
}
