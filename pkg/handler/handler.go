package handler

import (
	"github.com/gin-gonic/gin"
	"newProjectFolders/testTask1/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Routes() *gin.Engine {
	router := gin.New()
	task := router.Group("/task")
	{
		task.GET("/:id", h.getTaskById)
		task.POST("/", h.createTask)
		task.PUT("/:id", h.updateTaskById)
		task.DELETE("/:id", h.deleteTaskById)
	}
	return router
}
