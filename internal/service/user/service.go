package user

import (
	"github.com/biryanim/denet_tz/internal/client/db"
	"github.com/biryanim/denet_tz/internal/config"
	"github.com/biryanim/denet_tz/internal/repository"
	"github.com/biryanim/denet_tz/internal/service"
)

var _ service.UserService = (*serv)(nil)

type serv struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
	jwtConfig      config.JWTConfig
}

func NewService(userRepository repository.UserRepository, txManager db.TxManager) *serv {
	return &serv{
		userRepository: userRepository,
		txManager:      txManager,
	}
}
