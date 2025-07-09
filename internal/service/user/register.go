package user

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"

	apperrors "github.com/biryanim/denet_tz/internal/errors"
	"github.com/biryanim/denet_tz/internal/model"

	"github.com/pkg/errors"
)

func (s *serv) Register(ctx context.Context, userCreate *model.UserCreate) (int64, error) {
	existingUser, err := s.userRepository.GetByEmail(ctx, userCreate.Info.Email)
	if err != nil && !errors.Is(err, apperrors.ErrUserNotFound) {
		return 0, fmt.Errorf("failed to check existing user: %w", err)
	}
	if existingUser != nil {
		return 0, apperrors.ErrUserAlreadyExists
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userCreate.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("failed to hash password: %w", err)
	}

	userCreate.Password = string(passwordHash)

	id, err := s.userRepository.Create(ctx, userCreate)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	return id, nil
}
