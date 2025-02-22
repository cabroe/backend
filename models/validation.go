package models

import "gorm.io/gorm"

type Validator interface {
	Validate() error
}

func ValidateModel(m Validator) error {
	return m.Validate()
}

func WithTransaction(db *gorm.DB, fn func(*gorm.DB) error) error {
	tx := db.Begin()
	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
