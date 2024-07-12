package database

type FindOneOption struct {
	Where []WhereOption
	OrderOption
}

type FindManyOption struct {
	FindOneOption
	PaginationOption
}

type PaginationOption struct {
	Limit  int
	Offset int
}

type OrderOption struct {
	Field string
	Desc  bool
}

type DefaultOption struct {
	OrderField string
	Limit      int
	Offset     int
}

type WhereOption interface {
	GetQuery() string
	GetValues() []any
}
