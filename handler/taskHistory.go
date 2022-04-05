package handler

import (
	"net/http"
	"novocaine-dev/helper"
	"novocaine-dev/task"
	"novocaine-dev/taskHistory"
	"novocaine-dev/user"

	"github.com/gin-gonic/gin"
)

type taskHistoryHandler struct {
	service     taskHistory.Service
	taskService task.Service
	userService user.Service
}

func NewTaskHistoryHandler(service taskHistory.Service, taskService task.Service, userService user.Service) *taskHistoryHandler {
	return &taskHistoryHandler{service, taskService, userService}
}

func (h *taskHistoryHandler) CreateTaskHistory(c *gin.Context) {
	var input taskHistory.CreateTaskHistoryInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to create task history", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.UserDetails(input.UserId)
	if err != nil {
		response := helper.APIResponse("Failed to create task history", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.taskService.FindTaskById(input.TaskId)
	if err != nil {
		response := helper.APIResponse("Failed to create task history", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newTaskHistories, err := h.service.CreateTaskHistory(input)
	if err != nil {
		response := helper.APIResponse("Failed to create task history", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success create task histories", http.StatusOK, "success", taskHistory.FormatTaskHistory(newTaskHistories))
	c.JSON(http.StatusOK, response)
}
