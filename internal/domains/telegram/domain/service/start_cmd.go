package service

import (
	"trade-union-service/internal/domains/telegram/domain/entity"
)

func (s *Service) StartCommand(dto entity.StartCommandServiceDTO) error {
	if err := s.api.SendMessage(entity.SendMessageApiDTO{
		Text:    s.translate(startCommandTranslateKey, dto.Lang),
		ChatID:  dto.ChatID,
		Options: nil,
	}); err != nil {
		s.log.WithField(chatIDLoggingKey, dto.ChatID).
			Error("service: StartCommand - s.api.SendMessage error", err.Error())
	}

	return nil
}
