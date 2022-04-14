package handler

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"novocaine-dev/helper"
	"novocaine-dev/product"
	"novocaine-dev/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	service product.Service
}

func NewProductHandler(service product.Service) *productHandler {
	return &productHandler{service}
}

func (h *productHandler) GetProducts(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	products, err := h.service.GetProducts(id)
	if err != nil {
		response := helper.APIResponse("Error to get product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of products", http.StatusOK, "success", product.FormatProducts(products))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) GetProductById(c *gin.Context) {
	// var input product.GetProductDetailInput
	id := c.Param("id")
	product_id, err := strconv.Atoi(id)
	if err != nil {
		response := helper.APIResponse("Failed to get product detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := h.service.GetProductById(product_id)
	response := helper.APIResponse("Product Detail", http.StatusOK, "success", product)
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) CreateProduct(c *gin.Context) {
	var input product.CreateProductInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to create product", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User_id = currentUser

	newProduct, err := h.service.CreateProduct(input)
	if err != nil {
		response := helper.APIResponse("Failed to create product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success create product", http.StatusOK, "success", product.FormatProduct(newProduct))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) UpdateProduct(c *gin.Context) {
	var inputID product.GetProductDetailInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData product.CreateProductInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to update product", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedProduct, err := h.service.UpdateProduct(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success create product", http.StatusOK, "success", product.FormatProduct(updatedProduct))
	c.JSON(http.StatusOK, response)

}

func (h *productHandler) DeleteProduct(c *gin.Context) {
	var inputID product.GetProductDetailInput
	err := c.ShouldBindUri(&inputID)

	deletedProduct, err := h.service.DeleteProduct(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete product", http.StatusOK, "success", deletedProduct)
	c.JSON(http.StatusOK, response)

}

func (h *productHandler) UploadImage(c *gin.Context) {
	var inputID product.GetProductDetailInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	file, err := c.FormFile("image_url")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to update product", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	path := "images/" + file.Filename
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to update product", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.SaveImage(inputID.ID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to update product", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Success upload image", http.StatusBadRequest, "success", data)
	c.JSON(http.StatusBadRequest, response)
	return
}

func (h *productHandler) CreateProductBulk(c *gin.Context) {
	var input product.CreateProductInput
	file, _, err := c.Request.FormFile("excelfile")
	if err != nil {
		response := helper.APIResponse("Failed to create product", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User_id = currentUser

	fmt.Println(currentUser)

	reader := csv.NewReader(file)
	reader.LazyQuotes = true

	var line []string
	const row = 3
	for {
		//store acquired data for each line in the line
		line, err = reader.Read()
		if err != nil {
			break
		}

		if line[0] == "name" {
			continue
		}

		input.Name = line[0]
		input.Serial_number = line[1]
		price, err := strconv.Atoi(line[2])
		if err != nil {
			response := helper.APIResponse("Failed to create product", http.StatusBadRequest, "error", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		input.Price = price

		// Create Product in the background
		go h.service.CreateProduct(input)
	}

	response := helper.APIResponse("Success create product!", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
