package service

import (
	"trade-union-service/internal/domains/telegram/domain/entity"
)

func (s *Service) StartCommand(dto entity.StartCommandServiceDTO) error {
	_, err := s.api.SendMessage(s.translate(startCommandTranslateKey, dto.Lang), dto.ChatID, nil)
	if err != nil {
		s.log.WithField(chatIDLoggingKey, dto.ChatID).Error("service: StartCommand - s.api.SendMessage error", err.Error())
	}
	return nil
}
