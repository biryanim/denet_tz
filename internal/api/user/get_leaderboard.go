package user

import (
	"github.com/biryanim/denet_tz/internal/converter"
	apperrors "github.com/biryanim/denet_tz/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	topTen = 10
)

func (i *Implementation) GetLeaderboard(c *gin.Context) {
	users, err := i.userService.GetLeaderboard(c.Request.Context(), topTen)
	if err != nil {
		appErr := apperrors.FromError(err)
		c.JSON(appErr.StatusCode, gin.H{"error": appErr.Error()})
		return
	}

	c.JSON(http.StatusOK, converter.ToUsersListResp(users))
}
