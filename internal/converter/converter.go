package converter

import (
	"github.com/biryanim/denet_tz/internal/api/dto"
	"github.com/biryanim/denet_tz/internal/model"
	"time"
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

func FromUserTaskCompleteReq(req *dto.UserTaskComplete) *model.UserTask {
	return &model.UserTask{
		UserID: req.UserID,
		TaskId: req.TaskID,
	}
}

func ToUserTaskCompleteResp(user *model.User) *dto.User {
	var t time.Time
	if user.UpdatedAt.Valid {
		t = user.UpdatedAt.Time
	}
	return &dto.User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Points:    user.Points,
		CreatedAt: user.CreatedAt,
		UpdatedAt: t,
	}
}

func FromReferralReq(ref *dto.Referral) *model.Referrals {
	return &model.Referrals{
		ReferrerUserId: ref.ReferrerId,
		ReferredUserId: ref.ReferredId,
	}
}

func ToUsersListResp(users []*model.User) *dto.Users {
	var (
		dtoUser []*dto.User
	)
	for _, user := range users {
		var t time.Time
		if user.UpdatedAt.Valid {
			t = user.UpdatedAt.Time
		}

		u := &dto.User{
			ID:        user.ID,
			Email:     user.Email,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
			UpdatedAt: t,
			Points:    user.Points,
		}
		dtoUser = append(dtoUser, u)
	}

	return &dto.Users{
		Users: dtoUser,
	}
}

func ToStatusResp(status *model.Status) *dto.Status {
	if status == nil {
		return nil
	}

	var t time.Time
	if status.User != nil && status.User.UpdatedAt.Valid {
		t = status.User.UpdatedAt.Time
	}

	var tasks []*dto.Task
	if status.Task != nil {
		for _, ts := range status.Task {
			if ts == nil {
				continue
			}
			tasks = append(tasks, &dto.Task{
				ID:           ts.ID,
				Name:         ts.Name,
				Description:  ts.Description,
				PointsReward: ts.PointsReward,
				CreatedAt:    ts.CreatedAt,
			})
		}
	}

	dtoStatus := &dto.Status{
		User: &dto.User{
			ID:        status.User.ID,
			Username:  status.User.Username,
			Email:     status.User.Email,
			CreatedAt: status.User.CreatedAt,
			UpdatedAt: t,
			Points:    status.User.Points,
		},
		Referrers: &dto.Referrers{
			UserIds: status.Referrers.UserIds,
		},
		Tasks: &dto.CompletedTask{
			Tasks: tasks,
		},
	}

	return dtoStatus
}
