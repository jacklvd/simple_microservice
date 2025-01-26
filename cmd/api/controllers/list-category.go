package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacklvd/simple_microservice/internal/repositories"
	usecases "github.com/jacklvd/simple_microservice/internal/use-cases"
)

func ListCategory(ctx *gin.Context, repository repositories.IRepository) {

	useCase := usecases.NewListCategoryUseCase(repository)
	categories, err := useCase.Execute()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   err.Error(),
				"success": false,
			})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"success": true, "categories": categories})
}
