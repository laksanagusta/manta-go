package handler

import (
	"net/http"
	"novocaine-dev/helper"
	"novocaine-dev/transactionProduct"

	"github.com/gin-gonic/gin"
)

type transactionProductHandler struct {
	service transactionProduct.Service
}

func NewTransactionProductHandler(service transactionProduct.Service) *transactionProductHandler {
	return &transactionProductHandler{service}
}

func (h *transactionProductHandler) Delete(c *gin.Context) {
	var input transactionProduct.ID
	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to delete items", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	deleteTransactionProduct, err := h.service.Delete(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete items", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete items", http.StatusOK, "success", deleteTransactionProduct)
	c.JSON(http.StatusOK, response)
}
