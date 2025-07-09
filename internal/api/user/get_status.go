package user

import (
	"fmt"
	"github.com/biryanim/denet_tz/internal/converter"
	apperrors "github.com/biryanim/denet_tz/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (i *Implementation) GetStatus(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || userId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	status, err := i.userService.GetStatus(c.Request.Context(), userId)
	if err != nil {
		fmt.Println(err)
		appErr := apperrors.FromError(err)
		c.JSON(appErr.StatusCode, gin.H{"error": appErr.Error()})
		return
	}

	c.JSON(http.StatusOK, converter.ToStatusResp(status))
}
