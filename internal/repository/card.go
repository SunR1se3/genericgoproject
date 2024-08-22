package repository

import (
	"GenericProject/internal/domain"
	"github.com/jmoiron/sqlx"
)

type CardRepository struct {
	db *sqlx.DB
}

func NewCardRepository(db *sqlx.DB) *CardRepository {
	return &CardRepository{db: db}
}

func (c CardRepository) GetCardByName(name string) (*domain.Card, error) {
	return nil, nil
}
