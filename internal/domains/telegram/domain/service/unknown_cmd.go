package service

import (
	"trade-union-service/internal/domains/telegram/domain/entity"
)

func (s *Service) UnknownCommand(dto entity.UnknownCommandServiceDTO) error {
	if err := s.api.SendMessage(entity.SendMessageApiDTO{
		Text:    s.translate(unknownCommandTranslateKey, dto.Lang),
		ChatID:  dto.ChatID,
		Options: nil,
	}); err != nil {
		s.log.WithField(chatIDLoggingKey, dto.ChatID).
			Error("service: UnknownCommand - s.api.SendMessage error", err.Error())
	}

	return nil
}
