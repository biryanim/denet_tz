package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	Points    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserInfo struct {
	Username string
	Email    string
}

type UserCreate struct {
	Info     UserInfo
	Password string
}

type UserLogin struct {
	Email    string
	Password string
}

type UserClaims struct {
	jwt.StandardClaims
	Email string
}
