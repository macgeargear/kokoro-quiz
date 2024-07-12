package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SqlDatabase struct {
	*gorm.DB
}

func (d *SqlDatabase) Close() error {
	sql, err := d.DB.DB()

	if err != nil {
		return err
	}

	return sql.Close()
}

func (d *SqlDatabase) SelectedActor(tx *gorm.DB) *SqlDatabase {
	if tx != nil {
		return &SqlDatabase{tx}
	}

	return d
}

func (r *SqlDatabase) AddWhereClause(tx *gorm.DB, whereOptions ...WhereOption) *SqlDatabase {
	preCondition := r.SelectedActor(tx)
	for _, s := range whereOptions {
		preCondition.DB = r.DB.Where(s.GetQuery(), s.GetValues()...)
	}

	return &SqlDatabase{preCondition.DB}
}

func (r *SqlDatabase) AddPagination(tx *gorm.DB, option PaginationOption) *SqlDatabase {
	preCondition := r.SelectedActor(tx)

	if option.Limit == 0 {
		preCondition.DB = preCondition.Limit(-1)
	} else {
		preCondition.DB = preCondition.Limit(option.Limit)
	}

	if option.Offset == 0 {
		preCondition.DB = preCondition.Offset(-1)
	} else {
		preCondition.DB = preCondition.Offset(option.Offset)
	}

	return &SqlDatabase{preCondition.DB}
}

func (r *SqlDatabase) AddOrder(tx *gorm.DB, option OrderOption) *SqlDatabase {
	preCondition := r.SelectedActor(tx)

	if option.Desc && option.Field != "" {
		preCondition.DB = preCondition.Order(clause.OrderByColumn{Column: clause.Column{Name: option.Field}, Desc: option.Desc})
	}

	return &SqlDatabase{preCondition.DB}
}
