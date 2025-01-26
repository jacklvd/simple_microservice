package usecases

import (
	"github.com/jacklvd/simple_microservice/internal/entities"
	"github.com/jacklvd/simple_microservice/internal/repositories"
)

type createCategoryUseCase struct {
	// db
	repository repositories.IRepository
}

func NewCreateCategoryUseCase(repository repositories.IRepository) *createCategoryUseCase {
	return &createCategoryUseCase{repository: repository}
}

func (u *createCategoryUseCase) Execute(name string) error {
	// business logic here
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}
	// verify if category already exists
	categories, _ := u.repository.List()

	for _, c := range categories {
		if c.Name == name {
			return entities.ErrCategoryAlreadyExists()
		}
	}

	err = u.repository.Save(category)

	if err != nil {
		return err
	}

	return nil
}
