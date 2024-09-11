package domain

import (
	"github.com/google/uuid"
	"time"
)

type Card struct {
	Id          uuid.UUID  `json:"id" db:"id"`
	Number      string     `json:"number" db:"number"`
	SomeCount   int        `json:"someCount" db:"some_count"`
	IsActive    bool       `json:"isActive" db:"is_active"`
	Responsible *string    `json:"responsible" db:"responsible"`
	CardType    CardType   `json:"cardTypeId" db:"card_type_id"`
	CreatedAt   time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt   *time.Time `json:"updatedAt" db:"updated_at"`
}

type CardDTO struct {
	Id          uuid.UUID `json:"id"`
	Number      string    `json:"number"`
	SomeCount   int       `json:"someCount"`
	IsActive    bool      `json:"isActive"`
	Responsible *string   `json:"responsible"`
	CardType    string    `json:"cardType"`
}

type CardCreateForm struct {
	Name        string  `json:"name"`
	SomeCount   int     `json:"someCount" db:"some_count"`
	IsActive    bool    `json:"isActive" db:"is_active"`
	Responsible *string `json:"responsible" db:"responsible"`
}

type CardUpdateForm struct {
	Name        string  `json:"name"`
	SomeCount   int     `json:"someCount" db:"some_count"`
	IsActive    bool    `json:"isActive" db:"is_active"`
	Responsible *string `json:"responsible" db:"responsible"`
}

func (f *CardCreateForm) Prepare(m *Card) {
	m.Id = uuid.New()
	m.SomeCount = f.SomeCount
	m.IsActive = f.IsActive
	m.Responsible = f.Responsible
	m.CreatedAt = time.Now()
}

func (f *CardUpdateForm) Prepare(m *Card) {
	m.SomeCount = f.SomeCount
	m.IsActive = f.IsActive
	m.Responsible = f.Responsible
	t := time.Now()
	m.UpdatedAt = &t
}
