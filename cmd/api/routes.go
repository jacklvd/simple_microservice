package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jacklvd/simple_microservice/cmd/api/controllers"
	"github.com/jacklvd/simple_microservice/internal/repositories"
)

func CategoryRoutes(router *gin.Engine) {
	// Simple group: categories
	categoryRoutes := router.Group("/categories")

	inMemoryCategoryRepository := repositories.NewInMemoryCategoryRepository()

	categoryRoutes.GET("/", func(ctx *gin.Context) {
		controllers.ListCategory(ctx, inMemoryCategoryRepository)
	})
	categoryRoutes.POST("/", func(ctx *gin.Context) {
		controllers.CreateCategory(ctx, inMemoryCategoryRepository)
	})

}
