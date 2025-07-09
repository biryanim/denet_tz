package user

import (
	"github.com/biryanim/denet_tz/internal/client/db"
	"github.com/biryanim/denet_tz/internal/repository"
	"github.com/biryanim/denet_tz/internal/service"
)

var _ service.UserService = (*serv)(nil)

const (
	referFriendTask   = "refer_friend"
	joinTelegramTask  = "join_telegram"
	followTwitterTask = "follow_twitter"
)

type serv struct {
	userRepo  repository.UserRepository
	taskRepo  repository.TaskRepository
	txManager db.TxManager
}

func NewUserService(userRepo repository.UserRepository, taskRepo repository.TaskRepository, txManager db.TxManager) *serv {
	return &serv{
		userRepo:  userRepo,
		taskRepo:  taskRepo,
		txManager: txManager,
	}
}
