package auth

import (
	"github.com/biryanim/denet_tz/internal/api/dto"
	"github.com/biryanim/denet_tz/internal/converter"
	apperrors "github.com/biryanim/denet_tz/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (i *Implementation) Login(c *gin.Context) {
	var loginReq dto.UserLoginRequest

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := i.authService.Login(c.Request.Context(), converter.FromUserLoginReq(&loginReq))

	if err != nil {
		appErr := apperrors.FromError(err)
		c.JSON(appErr.StatusCode, gin.H{"error": appErr.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.Token{
		Token: token,
	})
}
