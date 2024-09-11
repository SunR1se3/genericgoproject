package crud

import (
	"GenericProject/internal/pkg/constants"
	"GenericProject/internal/pkg/generic_injector"
	"GenericProject/internal/pkg/utils"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"reflect"
	"strings"
)

type BaseRepository[T any] struct {
	db *sqlx.DB
}

func NewBaseRepository[T any](db *sqlx.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

func (br *BaseRepository[T]) Create(entity T) error {
	fields, values := _entityData(entity, []string{})
	tableName := generic_injector.GI.GetTableName(reflect.TypeOf(entity))
	sql := fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s)", tableName, strings.Join(fields, ","), _insertPlaceholders(values))
	_, err := br.db.Exec(sql, values...)
	return err
}

func (br *BaseRepository[T]) GetOne(id uuid.UUID) (*T, error) {
	result := new(T)
	tableName := generic_injector.GI.GetTableName(reflect.TypeOf(result).Elem())
	sql := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", tableName)
	err := br.db.Get(result, sql, id)
	return result, err
}

func (br *BaseRepository[T]) GetAll() ([]T, error) {
	var result []T
	var entity T
	tableName := generic_injector.GI.GetTableName(reflect.TypeOf(entity))
	sql := fmt.Sprintf("SELECT * FROM %s", tableName)
	err := br.db.Select(&result, sql)
	return result, err
}

func (br *BaseRepository[T]) Update(entity T, id uuid.UUID) error {
	tableName := generic_injector.GI.GetTableName(reflect.TypeOf(entity))
	placeholders, values := _updatePlaceholders(entity, constants.UpdateExcludeFields)
	sql := fmt.Sprintf("UPDATE %s SET %s WHERE id = '%s'", tableName, placeholders, id)
	_, err := br.db.Exec(sql, values...)
	return err
}

func (br *BaseRepository[T]) Delete(id uuid.UUID) error {
	var entity T
	tableName := generic_injector.GI.GetTableName(reflect.TypeOf(entity))
	sql := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tableName)
	_, err := br.db.Exec(sql, id)
	return err
}

func _entityData(entity any, excludeFields []string) ([]string, []any) {
	var fields []string
	var values []any

	index := reflect.ValueOf(entity).NumField()
	for x := 0; x < index; x++ {
		fieldDbTag := reflect.TypeOf(entity).Field(x).Tag.Get("db")
		if utils.InArray(fieldDbTag, excludeFields) {
			continue
		}
		fields = append(fields, fieldDbTag)
		values = append(values, reflect.ValueOf(entity).Field(x).Interface())
	}
	return fields, values
}

func _insertPlaceholders(values []any) string {
	placeholders := ""
	for i, _ := range values {
		if i != len(values)-1 {
			placeholders += fmt.Sprintf("$%d, ", i+1)
		} else {
			placeholders += fmt.Sprintf("$%d", i+1)
		}

	}
	return placeholders
}

func _updatePlaceholders(entity any, excludeFields []string) (string, []any) {
	fields, values := _entityData(entity, excludeFields)
	placeholders := ""
	for i, _ := range values {
		if i != len(values)-1 {
			placeholders += fmt.Sprintf(fields[i]+" = $%d, ", i+1)
		} else {
			placeholders += fmt.Sprintf(fields[i]+" = $%d", i+1)
		}

	}
	return placeholders, values
}
