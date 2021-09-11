package main

import (
	"log"
	"pasarwarga/article"
	"pasarwarga/category"
	"pasarwarga/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pasarwarga?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// Category
	categoryRepository := category.NewRepository(db)
	categoryService := category.NewService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	// Article
	articleRepository := article.NewRepository(db)
	articleService := article.NewService(articleRepository)
	articleHandler := handler.NewArticleHandler(articleService)

	router := gin.Default()
	api := router.Group("/api/v1")

	// Article
	api.POST("/newarticle", articleHandler.RegisterArticle)
	api.GET("/article", articleHandler.GetAllArticle)
	api.GET("/article/:id", articleHandler.ArticleGetByID)
	api.DELETE("/deletearticle/:id", articleHandler.DeletedArticle)
	api.PUT("/updatearticle/:id", articleHandler.NewArticlesUpdate)

	// Category
	api.POST("/newcategory", categoryHandler.RegisterCategory)
	api.GET("/category", categoryHandler.GetAllCategory)
	api.GET("/category/:id", categoryHandler.GetCategoryByID)
	api.DELETE("/deletecategory/:id", categoryHandler.DeletedCategory)
	api.PUT("/updatecategory/:id", categoryHandler.NewCategoryUpdate)

	router.Run()

}
