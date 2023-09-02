package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/stg35/avito_test/docs"
	"github.com/stg35/avito_test/internal/service"
	"github.com/stg35/avito_test/internal/worker"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Handler struct {
	service         *service.Service
	taskDistributor worker.TaskDistributor
}

func NewHandler(service *service.Service, taskDistributor worker.TaskDistributor) *Handler {
	return &Handler{
		service,
		taskDistributor,
	}
}

func (h *Handler) InitRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		segment := api.Group("/segment")
		{
			segment.POST("/", h.createSegment)
			segment.DELETE("/:id", h.deleteSegment)
		}

		user := api.Group("/user")
		{
			user.POST("/", h.createUser)
			user.PATCH("/addSegments", h.addSegments)
			user.PATCH("/deleteSegments", h.deleteSegments)
			user.GET("/showSegments/:id", h.showSegments)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func statusResponse(message string) gin.H {
	return gin.H{"message": message}
}
