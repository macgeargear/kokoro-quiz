package question

import (
	"database/sql"

	"github.com/macgeargear/kokoro-quiz/internal/database"
	"gorm.io/gorm"
)

type Repository interface {
	Create(params Question, tx *gorm.DB) error
	Update(options database.FindOneOption, params Question, tx *gorm.DB) (int64, error)
	Delete(params database.FindOneOption, tx *gorm.DB) (int64, error)
	FindMany(params database.FindManyOption, tx *gorm.DB) ([]Question, error)
	FindOne(params database.FindOneOption, tx *gorm.DB) (Question, error)
	Count(tx *gorm.DB) (int64, error)
	GetTableName() string
	Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
}
