package handler

import (
	"net/http"
	"pasarwarga/category"
	"pasarwarga/helper"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type categoryHandler struct {
	categoryService category.Service
}

func NewCategoryHandler(categoryService category.Service) *categoryHandler {
	return &categoryHandler{categoryService}
}

func (h *categoryHandler) RegisterCategory(c *gin.Context) {

	var input category.CategoryUserInput

	err := c.ShouldBind(&input)
	if err != nil {
		var errors []string

		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())
		}

		errorMessage := gin.H{"errors": errors}

		response := helper.APIresponse("Failed Created Category", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newCategory, err := h.categoryService.RegisterCategory(input)

	if err != nil {
		response := helper.APIresponse("Failed Created Category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := category.FormatCategory(newCategory)
	response := helper.APIresponse("Category has been created", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) GetAllCategory(c *gin.Context) {
	categoys, err := h.categoryService.GetAllCategory()
	if err != nil {
		response := helper.APIresponse("Failed Find Category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse("Category Find", http.StatusOK, "success", categoys)
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) GetCategoryByID(c *gin.Context) {
	var input category.GetCategoryDetail
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIresponse("Failed Find Category Detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	categoris, err := h.categoryService.GetCategoryByID(input)
	if err != nil {
		response := helper.APIresponse("Failed Find Category Detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse("Success Find Category Detail", http.StatusOK, "success", categoris)
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) NewCategoryUpdate(c *gin.Context) {
	var inputID category.GetCategoryDetail

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIresponse("Failed Update Category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData category.CategoryUserInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		response := helper.APIresponse("Failed Update Category ", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedCategori, err := h.categoryService.Update(inputID, inputData)
	if err != nil {
		response := helper.APIresponse("Failed Update Category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse("Success Update Category", http.StatusOK, "success", updatedCategori)
	c.JSON(http.StatusOK, response)
}

func (h *categoryHandler) DeletedCategory(c *gin.Context) {
	var inputID category.GetCategoryDetail

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIresponse("Failed Delete Category", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	deletedCategori, err := h.categoryService.Delete(inputID)
	if err != nil {
		response := helper.APIresponse("Failed Delete Category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse("Success Delete Category", http.StatusOK, "success", deletedCategori)
	c.JSON(http.StatusOK, response)

}
