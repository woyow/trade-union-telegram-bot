package service

import (
	"trade-union-service/internal/domains/telegram/domain/entity"

	"github.com/sirupsen/logrus"
)

func (s *Service) StartCommand(dto entity.StartCommandServiceDTO) error {
	if err := s.api.SendMessage(entity.SendMessageAPIDTO{
		Text:    s.translate(startCommandTranslateKey, dto.Lang),
		ChatID:  dto.ChatID,
		Options: nil,
	}); err != nil {
		s.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			domainLoggingKey: domainLoggingValue,
			layerLoggingKey:  layerLoggingValue,
		}).Error("StartCommand - s.api.SendMessage error: ", err.Error())
	}

	return nil
}
