package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Card
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Card: NewCardRepository(db),
	}
}

type Card interface {
}
