package model

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type User struct {
	ID        int64
	Username  string
	Email     string
	Password  string
	Points    int64
	CreatedAt time.Time
	UpdatedAt sql.NullTime
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

type UserTask struct {
	UserID int64
	TaskId int64
}

type Task struct {
	ID           int64
	Name         string
	Description  string
	PointsReward int64
	CreatedAt    time.Time
}

type Referrals struct {
	ReferrerUserId int64
	ReferredUserId int64
}

type Referrers struct {
	UserIds []int64
}

type Status struct {
	User      *User
	Referrers *Referrers
	Task      []*Task
}
