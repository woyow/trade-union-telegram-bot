package service

import (
	"context"

	"trade-union-service/internal/domains/users/domain/entity"

	"github.com/sirupsen/logrus"
)

type repo interface {
	CreateUser(ctx context.Context, dto entity.CreateUserRepoDTO) (*entity.CreateUserOut, error)
	GetUser(ctx context.Context, dto entity.GetUserRepoDTO) (*entity.GetUserOut, error)
	UpdateUser(ctx context.Context, dto entity.UpdateUserRepoDTO) error
}

type cache interface {
}

type Service struct {
	repo  repo
	cache cache
	log   *logrus.Logger
}

func NewService(repo repo, cache cache, log *logrus.Logger) *Service {
	return &Service{
		repo:  repo,
		cache: cache,
		log:   log,
	}
}
