package user

import (
	"context"
	"github.com/biryanim/denet_tz/internal/model"
)

func (s *serv) GetStatus(ctx context.Context, userId int64) (*model.Status, error) {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	completedTasks, err := s.taskRepo.GetUserCompletedTasks(ctx, userId)
	if err != nil {
		return nil, err
	}

	referrers, err := s.userRepo.GetReferrers(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &model.Status{
		User:      user,
		Task:      completedTasks,
		Referrers: referrers,
	}, nil
}
