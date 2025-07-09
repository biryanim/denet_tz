package user

import (
	"github.com/biryanim/denet_tz/internal/api/dto"
	"github.com/biryanim/denet_tz/internal/converter"
	apperrors "github.com/biryanim/denet_tz/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (i *Implementation) AddReferrer(c *gin.Context) {
	var (
		ref dto.Referral
		err error
	)
	ref.ReferredId, err = strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil || ref.ReferredId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err = c.ShouldBindJSON(&ref); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	success, err := i.userService.AddReferrer(c.Request.Context(), converter.FromReferralReq(&ref))
	if err != nil {
		appErr := apperrors.FromError(err)
		c.JSON(appErr.StatusCode, gin.H{"error": appErr.Error()})
		return
	}

	if !success {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add referrer"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"referrer": ref.ReferredId})
}
