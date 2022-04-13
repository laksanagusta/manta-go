package handler

import (
	"net/http"
	"novocaine-dev/helper"
	"novocaine-dev/transaction"
	"novocaine-dev/user"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) AddToCart(c *gin.Context) {
	var input transaction.AddToCartInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User_id = currentUser.ID

	newTransaction, err := h.service.AddToCart(input)
	if err != nil {
		response := helper.APIResponse("Failed to create transaction", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success create task", http.StatusOK, "success", transaction.FormatTransaction(newTransaction))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) FindById(c *gin.Context) {
	var input transaction.FindById
	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to load transaction", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loadTransaction, err := h.service.FindById(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to load transaction", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success load transaction", http.StatusOK, "success", transaction.FormatTransaction(loadTransaction))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) FindTransactionByUser(c *gin.Context) {
	var input transaction.FindByUserId
	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to load transaction", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loadTransaction, err := h.service.FindTransactionByUser(input.UserID)
	if err != nil {
		response := helper.APIResponse("Failed to load transaction", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success load transaction", http.StatusOK, "success", transaction.FormatTransactions(loadTransaction))
	c.JSON(http.StatusOK, response)
}
