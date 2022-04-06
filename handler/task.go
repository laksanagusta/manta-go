package handler

import (
	"net/http"
	"novocaine-dev/helper"
	"novocaine-dev/task"
	"novocaine-dev/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type taskHandler struct {
	service     task.Service
	userService user.Service
}

func NewTaskHandler(service task.Service, userService user.Service) *taskHandler {
	return &taskHandler{service, userService}
}

func (h *taskHandler) CreateTask(c *gin.Context) {
	var input task.CreateTaskInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to create task", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// findUserAssigne, err := h.userService.UserDetails(input.TaskAssignedTo)
	// if err != nil {
	// 	response := helper.APIResponse("Failed to create task, assigned user not found", http.StatusBadRequest, "error", err)
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// currentUser := c.MustGet("currentUser").(user.User)
	// input.TaskCreatedBy = currentUser
	// input.TaskAssignedTo = findUserAssigne.ID

	newTask, err := h.service.CreateTask(input)
	if err != nil {
		response := helper.APIResponse("Failed to create task", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success create task", http.StatusOK, "success", task.FormatTask(newTask))
	c.JSON(http.StatusOK, response)
}

func (h *taskHandler) UpdateTask(c *gin.Context) {
	var input task.UpdateTaskInput
	id, _ := strconv.Atoi(c.Param("id"))
	input.ID = id

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to create task", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// _, err = h.userService.UserDetails(input.TaskAssignedTo)
	// var message = "Failed to create task, assigned user not found"
	// if err != nil {
	// 	response := helper.APIResponse(message, http.StatusBadRequest, "error", err)
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	updateTask, err := h.service.UpdateTask(input)
	if err != nil {
		response := helper.APIResponse("Failed to update task", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success update task", http.StatusOK, "success", task.FormatTask(updateTask))
	c.JSON(http.StatusOK, response)
}

func (h *taskHandler) FindTaskById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	findTask, err := h.service.FindTaskById(id)
	if err != nil {
		response := helper.APIResponse("Failed to get task", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success get task data", http.StatusOK, "success", task.FormatTask(findTask))
	c.JSON(http.StatusOK, response)
}

func (h *taskHandler) CustomFilter(c *gin.Context) {
	query := c.Request.URL.Query()
	tasks, err := h.service.CustomFilter(query)
	if err != nil {
		response := helper.APIResponse("Failed to get task", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success get task data", http.StatusOK, "success", task.FormatTasks(tasks))
	c.JSON(http.StatusOK, response)
}

func (h *taskHandler) Delete(c *gin.Context) {
	var taskId task.FindTaskById
	err := c.ShouldBindUri(&taskId)

	if err != nil {
		response := helper.APIResponse("Failed to delete task", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	deleteTask, err := h.service.Delete(taskId)
	if err != nil {
		response := helper.APIResponse("Failed to delete task", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete product", http.StatusOK, "success", deleteTask)
	c.JSON(http.StatusOK, response)
}

func (h *taskHandler) ProcessTask(c *gin.Context) {
	var input task.ProcessTaskInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to process task", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	findUser, err := h.userService.UserDetails(input.UserID)
	if err != nil {
		response := helper.APIResponse("Failed to process task", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var username string = findUser.Username

	currentUser := c.MustGet("currentUser").(user.User)
	input.UserIdProcessed = currentUser.ID

	processTask, err := h.service.ProcessTask(input, username)
	if err != nil {
		response := helper.APIResponse("Failed to process task", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success process task", http.StatusOK, "success", processTask)
	c.JSON(http.StatusOK, response)
}

func (h *taskHandler) AssignTask(c *gin.Context) {
	var input task.AssignTaskInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to assign task", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.UserDetails(input.UserID)
	if err != nil {
		response := helper.APIResponse("Failed to assign task", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	findTask, err := h.service.FindTaskById(input.TaskID)
	if err != nil {
		response := helper.APIResponse("Failed to assign task", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.AssignTask(input)
	if err != nil {
		response := helper.APIResponse("Failed to assign task", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success assign task", http.StatusOK, "success", findTask)
	c.JSON(http.StatusOK, response)
}
