package question

import (
	"database/sql"

	"github.com/macgeargear/kokoro-quiz/internal/database"
	"github.com/macgeargear/kokoro-quiz/internal/utils"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	entity    *Question
	db        *database.SqlDatabase
	tableName string
}

func NewRepository(db *database.SqlDatabase) Repository {
	return &repositoryImpl{
		entity:    &Question{},
		db:        db,
		tableName: "questions",
	}
}

func (r *repositoryImpl) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		return fc(tx)
	}, opts...)
}

func (r *repositoryImpl) Create(params Question, tx *gorm.DB) error {
	question := &Question{
		Type:          params.Type,
		Text:          params.Text,
		Options:       params.Options,
		CorrectAnswer: params.CorrectAnswer,
	}

	actor := r.db.SelectedActor(tx)

	if result := actor.Model(r.entity).Create(&question); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repositoryImpl) Update(options database.FindOneOption, params Question, tx *gorm.DB) (int64, error) {
	result := r.db.AddWhereClause(tx, options.Where...).Model(r.entity).Updates(params)

	return result.RowsAffected, result.Error
}

func (r *repositoryImpl) Delete(params database.FindOneOption, tx *gorm.DB) (int64, error) {
	result := r.db.AddWhereClause(tx, params.Where...).
		Model(r.entity).
		Updates(Question{DeletedAt: gorm.DeletedAt{Time: utils.TimeUTC(nil)}})

	return result.RowsAffected, result.Error
}

func (r *repositoryImpl) FindMany(params database.FindManyOption, tx *gorm.DB) (questions []Question, err error) {
	result := r.db.AddWhereClause(tx, params.Where...).
		AddPagination(tx, params.PaginationOption).
		AddOrder(tx, params.OrderOption).
		Find(&questions)

	if result.Error != nil {
		return questions, result.Error
	}

	return questions, nil
}

func (r *repositoryImpl) FindOne(params database.FindOneOption, tx *gorm.DB) (question Question, err error) {
	result := r.db.AddWhereClause(tx, params.Where...).First(&question)

	if result.Error != nil {
		return question, result.Error
	}

	return question, nil
}

func (r *repositoryImpl) Count(tx *gorm.DB) (int64, error) {
	var count int64
	result := tx.Model(r.entity).Count(&count)

	if result.Error != nil {
		return 0, result.Error
	}

	return count, result.Error
}

func (r *repositoryImpl) GetTableName() string {
	return r.tableName
}
