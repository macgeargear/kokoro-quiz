package database

import (
	"fmt"
	"strings"
)

type groupCondition struct {
	options   []WhereOption
	operation string
}

func (s groupCondition) GetQuery() string {
	queries := make([]string, 0, len(s.options))

	for _, spec := range s.options {
		queries = append(queries, spec.GetQuery())
	}

	return strings.Join(queries, fmt.Sprintf(" %s ", s.operation))
}

func (s groupCondition) GetValues() []any {
	values := make([]any, 0)

	for _, spec := range s.options {
		values = append(values, spec.GetValues()...)
	}

	return values
}

func And(options ...WhereOption) WhereOption {
	return groupCondition{
		options:   options,
		operation: "AND",
	}
}

func Or(options ...WhereOption) WhereOption {
	return groupCondition{
		options:   options,
		operation: "OR",
	}
}
