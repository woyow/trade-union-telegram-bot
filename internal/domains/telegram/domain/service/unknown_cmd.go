package service

import (
	"trade-union-service/internal/domains/telegram/domain/entity"

	"github.com/sirupsen/logrus"
)

func (s *Service) UnknownCommand(dto entity.UnknownCommandServiceDTO) error {
	if err := s.api.SendMessage(entity.SendMessageAPIDTO{
		Text:    s.translate(unknownCommandTranslateKey, dto.Lang),
		ChatID:  dto.ChatID,
		Options: nil,
	}); err != nil {
		s.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			domainLoggingKey: domainLoggingValue,
			layerLoggingKey:  layerLoggingValue,
		}).Error("UnknownCommand - s.api.SendMessage error", err.Error())
	}

	return nil
}
