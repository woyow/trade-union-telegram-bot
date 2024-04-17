package service

import (
	"context"

	"trade-union-service/internal/domains/telegram/domain/entity"

	"github.com/sirupsen/logrus"
)

const (
	chatIDLoggingKey   = "chat_id"
	domainLoggingKey   = "domain"
	domainLoggingValue = "telegram"
	layerLoggingKey    = "layer"
	layerLoggingValue  = "service"
	markdownParseMode  = "MarkdownV2"
)

type api interface {
	SendMessage(dto entity.SendMessageAPIDTO) error
	SendMessageWithInlineKeyboard(dto entity.SendMessageWithInlineKeyboardAPIDTO) error
}

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

	GetAppealSubjects(ctx context.Context, dto entity.GetAppealSubjectsRepoDTO) (entity.GetAppealSubjectsRepoOut, error)
}

type Service struct {
	repo          repo
	translateDict translateMap
	api           api
	log           *logrus.Logger
}

func NewService(repo repo, api api, log *logrus.Logger) *Service {
	return &Service{
		repo:          repo,
		translateDict: getTranslateMap(),
		api:           api,
		log:           log,
	}
}
