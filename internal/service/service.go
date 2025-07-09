package service

import (
	"context"
	"github.com/biryanim/denet_tz/internal/model"
)

type AuthService interface {
	Register(ctx context.Context, userCreate *model.UserCreate) (int64, error)
	Login(ctx context.Context, userLogin *model.UserLogin) (string, error)
	Check(ctx context.Context, token string) (bool, error)
}

type UserService interface {
	CompleteTask(ctx context.Context, userTask *model.UserTask) (*model.User, error)
	AddReferrer(ctx context.Context, ref *model.Referrals) (bool, error)
	GetLeaderboard(ctx context.Context, limit int64) ([]*model.User, error)
	GetStatus(ctx context.Context, userId int64) (*model.Status, error)
}
