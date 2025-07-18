package app

import (
	"context"
	"github.com/biryanim/denet_tz/internal/api/auth"
	"github.com/biryanim/denet_tz/internal/api/user"
	"github.com/biryanim/denet_tz/internal/repository/task"
	"log"

	"github.com/biryanim/denet_tz/internal/client/db"
	"github.com/biryanim/denet_tz/internal/client/db/pg"
	"github.com/biryanim/denet_tz/internal/client/db/transaction"
	"github.com/biryanim/denet_tz/internal/config"
	"github.com/biryanim/denet_tz/internal/config/env"
	"github.com/biryanim/denet_tz/internal/repository"
	userRepo "github.com/biryanim/denet_tz/internal/repository/user"
	"github.com/biryanim/denet_tz/internal/service"
	authServ "github.com/biryanim/denet_tz/internal/service/auth"
	userServ "github.com/biryanim/denet_tz/internal/service/user"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	httpConfig config.HTTPConfig
	jwtConfig  config.JWTConfig

	dbClient  db.Client
	txManager db.TxManager

	userRepository repository.UserRepository
	taskRepository repository.TaskRepository

	authService service.AuthService
	userService service.UserService

	authImpl *auth.Implementation
	userImpl *user.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to load pg config: %v", err)
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to load http config: %v", err)
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) JWTConfig() config.JWTConfig {
	if s.jwtConfig == nil {
		cfg, err := env.NewJWTConfig()
		if err != nil {
			log.Fatalf("failed to load jwt config: %v", err)
		}

		s.jwtConfig = cfg
	}
	return s.jwtConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to init db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping db: %v", err)
		}

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepo.NewRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) TaskRepository(ctx context.Context) repository.TaskRepository {
	if s.taskRepository == nil {
		s.taskRepository = task.NewRepository(s.DBClient(ctx))
	}

	return s.taskRepository
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authServ.NewService(s.UserRepository(ctx), s.TxManager(ctx), s.JWTConfig())
	}

	return s.authService
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userServ.NewUserService(s.UserRepository(ctx), s.TaskRepository(ctx), s.TxManager(ctx))
	}

	return s.userService
}

func (s *serviceProvider) AuthImpl(ctx context.Context) *auth.Implementation {
	if s.authImpl == nil {
		s.authImpl = auth.NewImplementation(s.AuthService(ctx))
	}

	return s.authImpl
}

func (s *serviceProvider) UserImpl(ctx context.Context) *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService(ctx))
	}

	return s.userImpl
}
