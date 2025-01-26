package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacklvd/simple_microservice/internal/repositories"
	usecases "github.com/jacklvd/simple_microservice/internal/use-cases"
)

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(ctx *gin.Context, repository repositories.IRepository) {
	var request CreateCategoryRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   err.Error(),
				"success": false,
			})
		return
	}

	useCase := usecases.NewCreateCategoryUseCase(repository)
	err := useCase.Execute(request.Name)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   err.Error(),
				"success": false,
			})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"success": true})
}
