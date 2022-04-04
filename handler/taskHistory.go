package handler

import (
	"net/http"
	"novocaine-dev/helper"
	"novocaine-dev/taskHistory"
	"novocaine-dev/user"

	"github.com/gin-gonic/gin"
)

type taskHistoryHandler struct {
	service     taskHistory.Service
	userService user.Service
}

func NewTaskHistoryHandler(service taskHistory.Service, userService user.Service) *taskHistoryHandler {
	return &taskHistoryHandler{service, userService}
}

func (h *taskHistoryHandler) CreateTask(c *gin.Context) {
	var input taskHistory.CreateTaskHistoryInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to create task history", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// findUserAssigne, err := h.userService
	// if err != nil {
	// 	response := helper.APIResponse("Failed to create task, assigned user not found", http.StatusBadRequest, "error", err)
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// currentUser := c.MustGet("currentUser").(user.User)
	// input.TaskCreatedBy = currentUser
	// input.TaskAssignedTo = findUserAssigne.ID

	// newTask, err := h.service.CreateTask(input)
	// if err != nil {
	// 	response := helper.APIResponse("Failed to create task", http.StatusBadRequest, "error", err)
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// response := helper.APIResponse("Success create task", http.StatusOK, "success", task.FormatTask(newTask))
	// c.JSON(http.StatusOK, response)
}
