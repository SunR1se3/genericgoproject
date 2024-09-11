package domain

import "github.com/google/uuid"

type CardType struct {
	Id    uuid.UUID `json:"id" db:"id"`
	Title string    `json:"title" db:"title"`
}
