package handler

import (
	"net/http"
	"pasarwarga/article"
	"pasarwarga/helper"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type articleHandler struct {
	articleService article.Service
}

func NewArticleHandler(articleService article.Service) *articleHandler {
	return &articleHandler{articleService}
}

func (h *articleHandler) RegisterArticle(c *gin.Context) {

	var input article.ArticleUserInput

	err := c.ShouldBind(&input)
	if err != nil {
		var errors []string

		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())
		}

		errorMessage := gin.H{"errors": errors}

		response := helper.APIresponse("Failed Created article", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newArticle, err := h.articleService.RegisterArticle(input)

	if err != nil {
		response := helper.APIresponse("Failed Created article", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := article.FormatArticle(newArticle)
	response := helper.APIresponse("Article has been created", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *articleHandler) GetAllArticle(c *gin.Context) {
	articles, err := h.articleService.GetAllArticle()
	if err != nil {
		response := helper.APIresponse("Failed Find Article", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse("Article Find", http.StatusOK, "success", articles)
	c.JSON(http.StatusOK, response)
}

func (h *articleHandler) NewArticlesUpdate(c *gin.Context) {
	var inputID article.GetArticleDetail

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIresponse("Failed Update Article", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData article.ArticleUserInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		response := helper.APIresponse("Failed Update Article ", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedArticle, err := h.articleService.Update(inputID, inputData)
	if err != nil {
		response := helper.APIresponse("Failed Update Article", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse("Success Update Article", http.StatusOK, "success", updatedArticle)
	c.JSON(http.StatusOK, response)
}

func (h *articleHandler) DeletedArticle(c *gin.Context) {
	var inputID article.GetArticleDetail

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIresponse("Failed Delete Article", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	deletedArticle, err := h.articleService.Delete(inputID)
	if err != nil {
		response := helper.APIresponse("Failed Delete Article", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse("Success Delete Article", http.StatusOK, "success", deletedArticle)
	c.JSON(http.StatusOK, response)

}

func (h *articleHandler) ArticleGetByID(c *gin.Context) {
	var input article.GetArticleDetail

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIresponse("Failed Find Article Detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	articlesDetails, err := h.articleService.GetArticleByID(input)
	if err != nil {
		response := helper.APIresponse("Failed Find Article Detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse("Success Find Article Detail", http.StatusOK, "success", article.AllFormatarticle(articlesDetails))
	c.JSON(http.StatusOK, response)
}
