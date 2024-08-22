package domain

import "github.com/google/uuid"

type Card struct {
	Id          uuid.UUID `json:"id" db:"id"`
	Number      string    `json:"number" db:"number"`
	SomeCount   int       `json:"someCount" db:"some_count"`
	IsActive    bool      `json:"isActive" db:"is_active"`
	Responsible *string   `json:"responsible" db:"responsible"`
}
