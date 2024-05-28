package repo

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type ClassRepo struct {
	db *gorm.DB
}

func NewClassRepo(db *gorm.DB) ClassRepo {
	return ClassRepo{db}
}

func (c ClassRepo) Init(data []model.Class) error {
	for _, class := range data {
		err := c.db.Create(&class).Error
		if err != nil {
			return err
		}
	}
	return nil // TODO: replace this
}