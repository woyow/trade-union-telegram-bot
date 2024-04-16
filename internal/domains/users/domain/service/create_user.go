package service

import (
	"context"

	entity "trade-union-service/internal/domains/users/domain/entity"
)

func (s *Service) CreateUser(ctx context.Context, dto entity.CreateUserServiceDTO) (*entity.CreateUserOut, error) {
	out, err := s.repo.CreateUser(ctx, entity.CreateUserRepoDTO{
		Roles:    dto.Roles,
		Fname:    dto.Fname,
		Lname:    dto.Lname,
		Mname:    dto.Mname,
		Position: dto.Position,
		ChatID:   dto.ChatID,
	})
	if err != nil {
		s.log.Error("service: CreateUser - s.repo.CreateUser error: ", err.Error())

		return nil, err
	}

	return out, nil
}
