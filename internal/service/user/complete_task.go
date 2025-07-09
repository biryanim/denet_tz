package user

import (
	"context"
	apperrors "github.com/biryanim/denet_tz/internal/errors"
	"github.com/biryanim/denet_tz/internal/model"
)

func (s *serv) CompleteTask(ctx context.Context, userTask *model.UserTask) (*model.User, error) {
	completed, err := s.userRepo.IsTaskCompleted(ctx, userTask.UserID, userTask.TaskId)
	if err != nil {
		return nil, err
	}

	if completed {
		return nil, apperrors.ErrTaskAlreadyCompleted
	}

	task, err := s.taskRepo.GetByID(ctx, userTask.TaskId)
	if err != nil {
		return nil, err
	}

	var user *model.User
	err = s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		err = s.userRepo.CompleteTask(ctx, userTask.UserID, userTask.TaskId)
		if err != nil {
			return err
		}

		user, err = s.userRepo.UpdatePoints(ctx, userTask.UserID, task.PointsReward)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}
