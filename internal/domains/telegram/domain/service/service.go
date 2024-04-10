package service

import (
	"context"
	"github.com/NicoNex/echotron/v3"
	"github.com/sirupsen/logrus"
	"trade-union-service/internal/domains/telegram/domain/entity"
)

const (
	loggingKey = "layer"
	loggingValue = "service"
	chatIDLoggingKey = "chatID"

	markdownParseMode = "MarkdownV2"
)


type repo interface {
	// State management
	CreateChatCurrentState(ctx context.Context, dto entity.CreateChatCurrentStateRepoDTO) (*entity.CreateChatCurrentStateOut, error)
	SetChatCurrentState(ctx context.Context, dto entity.SetChatCurrentStateRepoDTO) error
	GetChatCurrentState(ctx context.Context, dto entity.GetChatCurrentStateRepoDTO) (*entity.GetChatCurrentStateOut, error)

	// Appeals
	DeleteDraftAppeal(ctx context.Context, dto entity.DeleteDraftAppealRepoDTO) error
	CreateAppeal(ctx context.Context, dto entity.CreateAppealRepoDTO) (*entity.CreateAppealOut, error)
	UpdateDraftAppeal(ctx context.Context, dto entity.UpdateAppealRepoDTO) error
	GetDraftAppeal(ctx context.Context, dto entity.GetDraftAppealRepoDTO) (*entity.GetDraftAppealRepoOut, error)
}

type translateMap map[string]map[string]string

type Service struct {
	repo  repo
	translate translateMap
	api   *echotron.API
	log   *logrus.Logger
}

func NewService(repo repo, api *echotron.API, log *logrus.Logger) *Service {
	return &Service{
		repo:  repo,
		translate: getTranslateMap(),
		api:   api,
		log:   log.WithField(loggingKey, loggingValue).Logger,
	}
}
