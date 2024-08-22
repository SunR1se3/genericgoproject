package service

import (
	"GenericProject/internal/domain"
	"GenericProject/internal/repository"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

var Services *Service

type Service struct {
	Card
}

func NewService(r *repository.Repository, db *sqlx.DB) *Service {
	return &Service{
		Card: NewCardService(r, db),
	}
}

type Card interface {
	CreateCard(data domain.Card) (*uuid.UUID, error)
	GetById(id uuid.UUID) (*domain.Card, error)
	UpdateCard(data domain.Card) error
	GetCardByName(name string) (*domain.Card, error)
}
