package crud

import "github.com/google/uuid"

type IBaseRepository[T any] interface {
	Create(entity T) error
	GetOne(id uuid.UUID) (*T, error)
	GetAll() ([]T, error)
	Update(entity T, id uuid.UUID) error
	Delete(id uuid.UUID) error
}
