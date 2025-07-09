package user

import (
	"context"

	apperrors "github.com/biryanim/denet_tz/internal/errors"
	"github.com/biryanim/denet_tz/internal/model"
)

func (s *serv) AddReferrer(ctx context.Context, ref *model.Referrals) (bool, error) {
	hasReferrer, err := s.userRepo.HasReferrer(ctx, ref.ReferredUserId)
	if err != nil {
		return false, err
	}
	if hasReferrer {
		return false, apperrors.ErrUserAlreadyHasReferrer
	}

	_, err = s.userRepo.GetByID(ctx, ref.ReferrerUserId)
	if err != nil {
		return false, apperrors.ErrUserNotFound
	}

	err = s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		err = s.userRepo.AddReferrer(ctx, ref)
		if err != nil {
			return err
		}
		task, err := s.taskRepo.GetTaskByName(ctx, referFriendTask)
		if err != nil {
			return err
		}

		err = s.userRepo.CompleteTask(ctx, ref.ReferrerUserId, task.ID)
		if err != nil {
			return err
		}

		_, err = s.userRepo.UpdatePoints(ctx, ref.ReferrerUserId, task.PointsReward)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return false, err
	}

	return true, nil
}
