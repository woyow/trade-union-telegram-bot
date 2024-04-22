package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"

	"trade-union-service/internal/domains/users/domain/entity"
	"trade-union-service/internal/domains/users/errs"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	defaultCapacity = 6 // Number of possible fields to update
)

func (r *RepoImpl) UpdateUser(ctx context.Context, dto entity.UpdateUserRepoDTO) error {
	objectID, err := primitive.ObjectIDFromHex(dto.ID)
	if err != nil {
		r.log.Error("mongo: UpdateUser - primitive.ObjectIDFromHex error: ", err.Error())
		return err
	}

	filter := bson.D{{
		Key:   "_id",
		Value: objectID,
	}}

	update := make(bson.D, 0, defaultCapacity)

	if dto.Fname != nil {
		update = append(update, bson.E{
			Key: "$set",
			Value: bson.D{{
				Key:   "fname",
				Value: *dto.Fname,
			}},
		})
	}

	if dto.Lname != nil {
		update = append(update, bson.E{
			Key: "$set",
			Value: bson.D{{
				Key:   "lname",
				Value: *dto.Lname,
			}},
		})
	}

	if dto.Mname != nil {
		update = append(update, bson.E{
			Key: "$set",
			Value: bson.D{{
				Key:   "mname",
				Value: *dto.Mname,
			}},
		})
	}

	if dto.Position != nil {
		update = append(update, bson.E{
			Key: "$set",
			Value: bson.D{{
				Key:   "position",
				Value: *dto.Position,
			}},
		})
	}

	if dto.ChatID != nil {
		update = append(update, bson.E{
			Key: "$set",
			Value: bson.D{{
				Key:   "chatId",
				Value: *dto.ChatID,
			}},
		})
	}

	if dto.Roles != nil {
		update = append(update, bson.E{
			Key: "$set",
			Value: bson.D{{
				Key:   "chatId",
				Value: dto.Roles,
			}},
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
