package service

import (
	"context"

	"trade-union-service/internal/domains/users/domain/entity"
)

func (s *Service) UpdateUser(ctx context.Context, dto entity.UpdateUserServiceDTO) (*entity.GetUserOut, error) {
	if err := s.repo.UpdateUser(ctx, entity.UpdateUserRepoDTO{
		ID:       dto.ID,
		Roles:    dto.Roles,
		Fname:    dto.Fname,
		Lname:    dto.Lname,
		Mname:    dto.Mname,
		Position: dto.Position,
		ChatID:   dto.ChatID,
	}); err != nil {
		s.log.Error("service: UpdateUser - s.repo.UpdateUser error: ", err.Error())
		return nil, err
	}

	out, err := s.repo.GetUser(ctx, entity.GetUserRepoDTO{
		ID:     &dto.ID,
		ChatID: nil,
	})
	if err != nil {
		s.log.Error("service: UpdateUser - s.repo.GetUser error: ", err.Error())
		return nil, err
	}

	return out, nil
}
