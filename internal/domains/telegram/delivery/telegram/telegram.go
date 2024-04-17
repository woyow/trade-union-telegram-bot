package telegram

import (
	"context"
	"github.com/NicoNex/echotron/v3"
	"github.com/sirupsen/logrus"
	"trade-union-service/internal/domains/telegram/domain/entity"
)

const (
	domainLoggingKey   = "domain"
	domainLoggingValue = "telegram"
	layerLoggingKey    = "layer"
	layerLoggingValue  = "telegram delivery"

	chatIDLoggingKey = "chat_id"
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
	NewCommandFullNameState(ctx context.Context, dto entity.NewCommandFullNameStateServiceDTO) error
	NewCommandSubjectState(ctx context.Context, dto entity.NewCommandSubjectStateServiceDTO) error
	NewCommandConfirmationState(ctx context.Context, dto entity.NewCommandConfirmationStateServiceDTO) error
}

type Telegram struct {
	dispatcher *echotron.Dispatcher
	log        *logrus.Logger
}

type destructChatID int64

func NewTelegram(service service, token string, log *logrus.Logger) *Telegram {
	destructCh := make(chan destructChatID)
	dispatcher := echotron.NewDispatcher(token, newBot(destructCh, service, log))

	go destructBot(destructCh, dispatcher, log)

	return &Telegram{
		dispatcher: dispatcher,
		log:        log,
	}
}

func destructBot(destructCh <-chan destructChatID, dispatcher *echotron.Dispatcher, log *logrus.Logger) {
	for {
		select {
		case b := <-destructCh:
			log.WithFields(logrus.Fields{
				domainLoggingKey: domainLoggingValue,
				layerLoggingKey:  layerLoggingValue,
			}).Info("destructBot - handle destruct for chatID: ", int64(b))

			dispatcher.DelSession(int64(b))

			log.WithFields(logrus.Fields{
				domainLoggingKey: domainLoggingValue,
				layerLoggingKey:  layerLoggingValue,
			}).Info("destructBot - destruct bot with chatID: ", int64(b))
		}
	}
}

func (b *Telegram) Run() error {
	defer b.log.WithFields(logrus.Fields{
		domainLoggingKey: domainLoggingValue,
		layerLoggingKey:  layerLoggingValue,
	}).Error("Run - stop telegram bot")

	b.log.WithFields(logrus.Fields{
		domainLoggingKey: domainLoggingValue,
		layerLoggingKey:  layerLoggingValue,
	}).Info("Run - start telegram bot")

	return b.dispatcher.Poll()
}
