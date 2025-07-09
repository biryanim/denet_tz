package auth

import (
	"context"
	"fmt"
	apperrors "github.com/biryanim/denet_tz/internal/errors"
	"github.com/biryanim/denet_tz/internal/model"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func (s *serv) Login(ctx context.Context, userLogin *model.UserLogin) (string, error) {
	user, err := s.userRepository.GetByEmail(ctx, userLogin.Email)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserNotFound) {
			return "", apperrors.ErrInvalidCredentials
		}
		return "", fmt.Errorf("failed to get auth: %w", err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password)); err != nil {
		return "", apperrors.ErrInvalidCredentials
	}

	token, err := generateToken(userLogin.Email, s.jwtConfig.TokenSecret(), s.jwtConfig.TokenExpiration())
	if err != nil {
		return "", errors.Wrap(err, "failed to generate token")
	}

	return token, nil
}
