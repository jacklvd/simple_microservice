package repositories

import "github.com/jacklvd/simple_microservice/internal/entities"

type IRepository interface {
	Save(category *entities.Category) error
	List() ([]*entities.Category, error)
}
