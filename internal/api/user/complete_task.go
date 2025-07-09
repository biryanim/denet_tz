package user

import (
	"fmt"
	"github.com/biryanim/denet_tz/internal/api/dto"
	"github.com/biryanim/denet_tz/internal/converter"
	apperrors "github.com/biryanim/denet_tz/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (i *Implementation) CompleteTask(c *gin.Context) {
	var (
		req dto.UserTaskComplete
		err error
	)
	req.UserID, err = strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil || req.UserID < 1 {
		fmt.Println("bbbbbbb", req.UserID)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err = c.ShouldBindJSON(&req); err != nil {
		fmt.Println("aaaaaaaaa")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	user, err := i.userService.CompleteTask(c.Request.Context(), converter.FromUserTaskCompleteReq(&req))
	if err != nil {
		appErr := apperrors.FromError(err)
		c.JSON(appErr.StatusCode, gin.H{"error": appErr.Error()})
		return
	}

	c.JSON(http.StatusOK, converter.ToUserTaskCompleteResp(user))
}
