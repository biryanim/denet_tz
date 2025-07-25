package env

import (
	"fmt"
	"os"
	"time"

	"github.com/biryanim/denet_tz/internal/config"

	"github.com/pkg/errors"
)

const (
	jwtTokenSecretKey  = "JWT_SECRET_KEY"
	jwtTokenExpiration = "JWT_TOKEN_EXPIRATION"
)

type jwtConfig struct {
	tokenSecret []byte
	tokenExp    time.Duration
}

func NewJWTConfig() (config.JWTConfig, error) {
	tokenSecret := []byte(os.Getenv(jwtTokenSecretKey))
	fmt.Println(tokenSecret)
	if len(tokenSecret) == 0 {
		return nil, errors.New("missing JWT token secret")
	}

	tokenExp, err := time.ParseDuration(os.Getenv(jwtTokenExpiration))
	fmt.Println(tokenSecret)
	if err != nil {
		return nil, err
	}

	return &jwtConfig{
		tokenSecret: tokenSecret,
		tokenExp:    tokenExp,
	}, nil
}

func (j *jwtConfig) TokenSecret() []byte {
	return j.tokenSecret
}

func (j *jwtConfig) TokenExpiration() time.Duration {
	return j.tokenExp
}
