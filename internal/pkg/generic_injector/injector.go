package generic_injector

import (
	"reflect"
)

var GI *GenericInjector

type GenericInjector struct {
	data map[string]any
}

func NewInjector() *GenericInjector {
	return &GenericInjector{
		data: make(map[string]any),
	}
}

func (m *GenericInjector) InjectModels(models ...reflect.Type) {
	for _, model := range models {
		m.data[model.Name()] = nil
		m.SetTableName(model, model.Name())
	}

}

func (m *GenericInjector) SetTableName(model reflect.Type, tableName string) {
	m.data[model.Name()] = map[string]any{
		"tableName": tableName,
	}
}

func (m *GenericInjector) GetTableName(model reflect.Type) string {
	return m.data[model.Name()].(map[string]any)["tableName"].(string)
}
