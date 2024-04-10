package telegram

import (
	"context"
	"github.com/NicoNex/echotron/v3"
	"github.com/sirupsen/logrus"
	"trade-union-service/internal/domains/telegram/domain/entity"
)

const (
	chatIDLoggingKey = "chatID"
)

type service interface {
	// State management
	CreateChatCurrentState(ctx context.Context, dto entity.CreateChatCurrentStateServiceDTO) error
	SetChatCurrentState(ctx context.Context, dto entity.SetChatCurrentStateServiceDTO) error
	GetChatCurrentState(ctx context.Context, dto entity.GetChatCurrentStateServiceDTO) (*entity.GetChatCurrentStateOut, error)

	// Unknown command service
	UnknownCommand(dto entity.UnknownCommandServiceDTO) error

	// Start command service
	StartCommand(dto entity.StartCommandServiceDTO) error

	// New command service
	NewCommand(ctx context.Context, dto entity.NewCommandServiceDTO) error
	NewCommandFirstNameState(ctx context.Context, dto entity.NewCommandFirstNameStateServiceDTO) error
	NewCommandLastNameState(ctx context.Context, dto entity.NewCommandLastNameStateServiceDTO) error
	NewCommandMiddleNameState(ctx context.Context, dto entity.NewCommandMiddleNameStateServiceDTO) error
	NewCommandConfirmationState(ctx context.Context, dto entity.NewCommandConfirmationStateServiceDTO) error
}

type Bot struct {
	dispatcher *echotron.Dispatcher
	log        *logrus.Logger
}

func NewBot(service service, token string, log *logrus.Logger) *Bot {
	dispatcher := echotron.NewDispatcher(token, newBot(service, log))
	return &Bot{
		dispatcher: dispatcher,
		log: log,
	}
}

func (b *Bot) Run() error {
	defer b.log.Error("bot: Run - stop telegram bot")
	b.log.Info("bot: Run - start telegram bot")
	return b.dispatcher.Poll()
}
