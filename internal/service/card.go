package service

import (
	"GenericProject/internal/domain"
	"GenericProject/internal/pkg/crud"
	"GenericProject/internal/pkg/mapper"
	"GenericProject/internal/repository"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"reflect"
)

type CardService struct {
	cardRepo repository.Card
	baseRepo crud.IBaseRepository[domain.Card]
}

func NewCardService(r repository.Card, db *sqlx.DB) *CardService {
	mapper.Map.SetTableName(reflect.TypeOf(domain.Card{}), "cards")
	baseRepo := crud.NewBaseRepository[domain.Card](db)
	return &CardService{cardRepo: r, baseRepo: baseRepo}
}

func (s *CardService) CreateCard(data domain.Card) (*uuid.UUID, error) {
	err := s.baseRepo.Create(data)
	return &data.Id, err
}

func (s *CardService) GetById(id uuid.UUID) (*domain.Card, error) {
	data, err := s.baseRepo.GetOne(id)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (s *CardService) UpdateCard(data domain.Card) error {
	return s.baseRepo.Update(data)
}

func (s *CardService) GetCardByName(name string) (*domain.Card, error) {
	return nil, nil
}
