package repository

import (
	"context"
	"github.com/biryanim/denet_tz/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.UserCreate) (int64, error)
	GetByID(ctx context.Context, id int64) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}
