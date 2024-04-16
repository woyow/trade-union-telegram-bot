package service

import (
	"context"

	"trade-union-service/internal/domains/telegram/domain/entity"

	"github.com/sirupsen/logrus"
)

func (s *Service) CreateChatCurrentState(ctx context.Context, dto entity.CreateChatCurrentStateServiceDTO) error {
	out, err := s.repo.CreateChatCurrentState(ctx, entity.CreateChatCurrentStateRepoDTO{
		State:  dto.State,
		ChatID: dto.ChatID,
	})
	if err != nil {
		s.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			domainLoggingKey: domainLoggingValue,
			layerLoggingKey:  layerLoggingValue,
		}).Error("CreateChatCurrentState - s.repo.CreateChatCurrentState error: ", err.Error())

		return nil
	}

	s.log.WithFields(logrus.Fields{
		chatIDLoggingKey: dto.ChatID,
		domainLoggingKey: domainLoggingValue,
	}).Info("CreateChatCurrentState - create chat state with id: ", out.ID)

	return nil
}

func (s *Service) SetChatCurrentState(ctx context.Context, dto entity.SetChatCurrentStateServiceDTO) error {
	if err := s.repo.SetChatCurrentState(ctx, entity.SetChatCurrentStateRepoDTO{
		State:  dto.State,
		ChatID: dto.ChatID,
	}); err != nil {
		s.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			domainLoggingKey: domainLoggingValue,
			layerLoggingKey:  layerLoggingValue,
		}).Error("SetChatCurrentState - s.repo.SetChatCurrentState error: ", err.Error())

		return nil
	}

	return nil
}

func (s *Service) GetChatCurrentState(ctx context.Context, dto entity.GetChatCurrentStateServiceDTO) (*entity.GetChatCurrentStateOut, error) {
	out, err := s.repo.GetChatCurrentState(ctx, entity.GetChatCurrentStateRepoDTO{
		ChatID: dto.ChatID,
	})
	if err != nil {
		s.log.WithFields(logrus.Fields{
			chatIDLoggingKey: dto.ChatID,
			domainLoggingKey: domainLoggingValue,
			layerLoggingKey:  layerLoggingValue,
		}).Error("GetChatCurrentState - s.repo.GetChatCurrentState error: ", err.Error())

		return nil, err
	}

	return out, nil
}
