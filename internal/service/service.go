package service

import (
	"context"
	"github.com/biryanim/denet_tz/internal/model"
)

type UserService interface {
	Register(ctx context.Context, userCreate *model.UserCreate) (int64, error)
	Login(ctx context.Context, userLogin *model.UserLogin) (string, error)
}
