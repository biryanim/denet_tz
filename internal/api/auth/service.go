package auth

import "github.com/biryanim/denet_tz/internal/service"

type Implementation struct {
	authService service.AuthService
}

func NewImplementation(authService service.AuthService) *Implementation {
	return &Implementation{
		authService: authService,
	}
}
