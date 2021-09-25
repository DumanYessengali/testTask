package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	tasks "newProjectFolders/testTask1"
	"strconv"
)

func (h *Handler) getTaskById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "INVALID_ID_PARAM")
		return
	}
	task, err := h.services.TaskInterface.GetTaskById(id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *Handler) createTask(c *gin.Context) {
	var input tasks.Tasks
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	task, err := h.services.TaskInterface.CreateTask(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

type UpdateTaskInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Author      *string `json:"author"`
}

func (h *Handler) updateTaskById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "INVALID_ID_PARAM")
		return
	}
	var input UpdateTaskInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	task, err := h.services.TaskInterface.GetTaskById(id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.TaskInterface.UpdateTaskById(id, input.Name, input.Description, input.Author, task); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Task updated")
}

func (h *Handler) deleteTaskById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "INVALID_ID_PARAM")
		return
	}

	if err := h.services.TaskInterface.DeleteTaskById(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Task Deleted")
}
