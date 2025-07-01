package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/quanergyO/ab_test_platform/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	apiV1 := router.Group("/api/v1")
	{
		abtest := apiV1.Group("/abtest")
		{
			abtest.POST("/", h.CreateABTest)
		}
	}

	return router
}
