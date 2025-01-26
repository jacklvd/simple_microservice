package usecases

import (
	"github.com/jacklvd/simple_microservice/internal/entities"
	"github.com/jacklvd/simple_microservice/internal/repositories"
)

type listCategoryUseCase struct {
	// db
	repository repositories.IRepository
}

func NewListCategoryUseCase(repository repositories.IRepository) *listCategoryUseCase {
	return &listCategoryUseCase{repository: repository}
}

func (u *listCategoryUseCase) Execute() ([]*entities.Category, error) {
	// business logic here
	category, err := u.repository.List()

	if err != nil {
		return nil, err
	}

	return category, nil
}
