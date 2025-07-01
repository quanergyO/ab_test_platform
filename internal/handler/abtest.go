package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quanergyO/ab_test_platform/internal/models"
)

func (h *Handler) CreateABTest(c *gin.Context) {
	var input models.ABTestSpec
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.service.GeneratePlugin(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	// TODO send sharedObject to service

	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "OK",
	})
}
