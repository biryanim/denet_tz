package repository

import (
	"context"
	"github.com/biryanim/denet_tz/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.UserCreate) (int64, error)
	GetByID(ctx context.Context, id int64) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	List(ctx context.Context, limit int64) ([]*model.User, error)
	UpdatePoints(ctx context.Context, id int64, points int64) (*model.User, error)
	IsTaskCompleted(ctx context.Context, userId int64, taskId int64) (bool, error)
	CompleteTask(ctx context.Context, userId int64, taskId int64) error
	HasReferrer(ctx context.Context, userId int64) (bool, error)
	AddReferrer(ctx context.Context, referrer *model.Referrals) error
	GetReferrers(ctx context.Context, userId int64) (*model.Referrers, error)
}

type TaskRepository interface {
	GetByID(ctx context.Context, id int64) (*model.Task, error)
	GetTaskByName(ctx context.Context, name string) (*model.Task, error)
	GetUserCompletedTasks(ctx context.Context, userId int64) ([]*model.Task, error)
}
