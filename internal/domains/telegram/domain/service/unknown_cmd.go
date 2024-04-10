package service

import (
	"trade-union-service/internal/domains/telegram/domain/entity"
)

const (
	unknownCommandMessage = "Неизвестная команда, попробуйте снова!"
)

func (s *Service) UnknownCommand(dto entity.UnknownCommandServiceDTO) error {
	_, err := s.api.SendMessage(s.translate[unknownCommandTranslateKey][dto.Lang], dto.ChatID, nil)
	if err != nil {
		s.log.WithField(chatIDLoggingKey, dto.ChatID).Error("service: UnknownCommand - s.api.SendMessage error", err.Error())
	}
	return nil
}
