package auth

import (
	"github.com/biryanim/denet_tz/internal/api/dto"
	"github.com/biryanim/denet_tz/internal/converter"
	apperrors "github.com/biryanim/denet_tz/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (i *Implementation) Register(c *gin.Context) {
	var registerReq dto.UserRegisterRequest

	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	resp, err := i.authService.Register(c.Request.Context(), converter.FromUserCreateReq(&registerReq))
	if err != nil {
		appErr := apperrors.FromError(err)
		c.JSON(appErr.StatusCode, gin.H{"error": appErr.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.UserRegisterResponse{ID: resp})
}
