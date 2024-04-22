package service

import (
	"trade-union-service/internal/domains/telegram/domain/entity"

	"github.com/sirupsen/logrus"
)

type sendSomethingWentWrongDTO struct {
	lang   string
	chatID int64
}

func (s *Service) sendSomethingWentWrong(dto sendSomethingWentWrongDTO) {
	if err := s.api.SendMessage(entity.SendMessageAPIDTO{
		Text:    s.translate(somethingWentWrongTranslateKey, dto.lang),
		ChatID:  dto.chatID,
		Options: nil,
	}); err != nil {
		s.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.chatID,
			domainLoggingKey: domainLoggingValue,
			layerLoggingKey:  layerLoggingValue,
		}).Error("sendSomethingWentWrong - s.api.SendMessage error: ", err.Error())
	}
}
