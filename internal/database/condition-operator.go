package database

import "fmt"

type conditionOption[T any] struct {
	field    string
	operator string
	value    T
}

func (s conditionOption[T]) GetQuery() string {
	return fmt.Sprintf("%s %s ?", s.field, s.operator)
}

func (s conditionOption[T]) GetValues() []any {
	return []any{s.value}
}

func Equal[T any](field string, value T) WhereOption {
	return conditionOption[T]{
		field:    field,
		operator: "=",
		value:    value,
	}
}

func GreaterThan[T comparable](field string, value T) WhereOption {
	return conditionOption[T]{
		field:    field,
		operator: ">",
		value:    value,
	}
}

func GreaterOrEqual[T comparable](field string, value T) WhereOption {
	return conditionOption[T]{
		field:    field,
		operator: ">=",
		value:    value,
	}
}

func LessThan[T comparable](field string, value T) WhereOption {
	return conditionOption[T]{
		field:    field,
		operator: "<",
		value:    value,
	}
}

func LessOrEqual[T comparable](field string, value T) WhereOption {
	return conditionOption[T]{
		field:    field,
		operator: ">=",
		value:    value,
	}
}

func In[T any](field string, value []T) WhereOption {
	return conditionOption[[]T]{
		field:    field,
		operator: "IN",
		value:    value,
	}
}
