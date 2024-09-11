package app_mapper

import (
	"GenericProject/internal/domain"
	"GenericProject/internal/pkg/mapper"
)

type AppMapper struct {
	CardMapper mapper.Mapper[domain.Card, domain.CardDTO]
}

var Mapper AppMapper

func NewMapper() AppMapper {
	cardMapper := mapper.Mapper[domain.Card, domain.CardDTO]{}
	cardMapper.AddMapping(func(src *domain.Card, dest *domain.CardDTO) {
		dest.Id = src.Id
		dest.Number = src.Number
		dest.Responsible = src.Responsible
		dest.SomeCount = src.SomeCount
		dest.CardType = src.CardType.Title
	})
	Mapper.CardMapper = cardMapper
	return Mapper
}
