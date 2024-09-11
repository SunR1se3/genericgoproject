package mapper

type MapperFunc[T1, T2 any] func(src *T1, dest *T2)

// Mapper — структура для хранения мапперов
type Mapper[T1, T2 any] struct {
	mappings []MapperFunc[T1, T2]
}

// AddMapping — добавляет функцию маппинга полей
func (m *Mapper[T1, T2]) AddMapping(mapping MapperFunc[T1, T2]) {
	m.mappings = append(m.mappings, mapping)
}

// Map — выполняет маппинг для переданных структур
func (m *Mapper[T1, T2]) Map(src *T1, dest *T2) {
	for _, mapping := range m.mappings {
		mapping(src, dest)
	}
}
