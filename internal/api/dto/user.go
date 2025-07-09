package dto

import (
	"time"
)

type UserRegisterRequest struct {
	Email    string `json:"email" binding:"required" validate:"required,email"`
	Username string `json:"username" binding:"required" validate:"required,min=2,max=32"`
	Password string `json:"password" binding:"required" validate:"required,min=8,max=32"`
}

type UserRegisterResponse struct {
	ID int64 `json:"id"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required" validate:"required,email"`
	Password string `json:"password" binding:"required" validate:"required,min=8,max=32"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type Token struct {
	Token string `json:"token"`
}

type UserTaskComplete struct {
	UserID int64
	TaskID int64 `json:"task_id" binding:"required"`
}

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Points    int64     `json:"points"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Referral struct {
	ReferrerId int64 `json:"referrer_id" binding:"required"`
	ReferredId int64 `json:"referred_id"`
}

type Users struct {
	Users []*User `json:"users"`
}

type Referrers struct {
	UserIds []int64 `json:"user_ids"`
}

type Task struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	PointsReward int64     `json:"points_reward"`
	CreatedAt    time.Time `json:"created_at"`
}

type CompletedTask struct {
	Tasks []*Task
}
type Status struct {
	User      *User          `json:"user"`
	Referrers *Referrers     `json:"referrers"`
	Tasks     *CompletedTask `json:"completed_tasks"`
}
