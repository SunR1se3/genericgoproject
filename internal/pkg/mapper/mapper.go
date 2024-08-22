package mapper

import (
	"reflect"
)

var Map *Mapper

type Mapper struct {
	data map[string]any
}

func NewMapper() *Mapper {
	return &Mapper{
		data: make(map[string]any),
	}
}

func (m *Mapper) InjectModels(models ...reflect.Type) {
	for _, model := range models {
		m.data[model.Name()] = nil
		m.SetTableName(model, model.Name())
	}

}

func (m *Mapper) SetTableName(model reflect.Type, tableName string) {
	m.data[model.Name()] = map[string]any{
		"tableName": tableName,
	}
}

func (m *Mapper) GetTableName(model reflect.Type) string {
	w := model.Name()
	k := m.data[w]
	q := k.(map[string]any)
	t := q["tableName"].(string)
	return t
	//return m.data[model.Name()].(map[string]any)["tableName"].(string)
}
