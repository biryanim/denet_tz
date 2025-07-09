package user

import (
	"context"
	"github.com/biryanim/denet_tz/internal/model"
)

func (s *serv) GetLeaderboard(ctx context.Context, limit int64) ([]*model.User, error) {
	users, err := s.userRepo.List(ctx, limit)
	if err != nil {
		return nil, err
	}

	return users, nil
}
