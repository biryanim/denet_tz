package user

import (
	"github.com/biryanim/denet_tz/internal/api/dto"
	"github.com/biryanim/denet_tz/internal/converter"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (i *Implementation) Register(c *gin.Context) {
	var registerReq dto.UserRegisterRequest

	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	resp, err := i.userService.Register(c.Request.Context(), converter.FromUserCreateReq(&registerReq))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, dto.UserRegisterResponse{ID: resp})
}
