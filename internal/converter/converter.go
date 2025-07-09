package converter

import (
	"github.com/biryanim/denet_tz/internal/api/dto"
	"github.com/biryanim/denet_tz/internal/model"
)

func FromUserCreateReq(createReq *dto.UserRegisterRequest) *model.UserCreate {
	return &model.UserCreate{
		Info: model.UserInfo{
			Username: createReq.Username,
			Email:    createReq.Email,
		},
		Password: createReq.Password,
	}
}

func FromUserLoginReq(loginReq *dto.UserLoginRequest) *model.UserLogin {
	return &model.UserLogin{
		Email:    loginReq.Email,
		Password: loginReq.Password,
	}
}
