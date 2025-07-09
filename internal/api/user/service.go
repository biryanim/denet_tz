package user

import "github.com/biryanim/denet_tz/internal/service"

type Implementation struct {
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
