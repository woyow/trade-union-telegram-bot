package service

import (
	"context"
	"trade-union-service/internal/domains/users/domain/entity"
)

func (s *Service) GetUser(ctx context.Context, dto entity.GetUserServiceDTO) (*entity.GetUserOut, error) {
	out, err := s.repo.GetUser(ctx, entity.GetUserRepoDTO{
		ID:     dto.ID,
		ChatID: dto.ChatID,
	})
	if err != nil {
		s.log.Error("service: GetUser - s.repo.GetUser error: ", err.Error())

		return nil, err
	}

	return out, nil
}
