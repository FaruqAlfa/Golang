package repo

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TeacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) TeacherRepo {
	return TeacherRepo{db}
}

func (t TeacherRepo) Save(data model.Teacher) error {
	return t.db.Create(&data).Error
 // TODO: replace this
}

func (t TeacherRepo) Query() ([]model.Teacher, error) {
	rows, err := t.db.Table("teachers").Select("*").Rows()
	if err != nil {
		return nil, err
	}
	var teachers []model.Teacher
	for rows.Next() {
		t.db.ScanRows(rows, &teachers)
	}

	return teachers, nil
}

func (t TeacherRepo) Update(id uint, name string) error {
	var teacher model.Teacher
	t.db.Model(&teacher).Where("id = ?", id).Update("name", name)
	return nil // TODO: replace this
}

func (t TeacherRepo) Delete(id uint) error {
	teacher := model.Teacher{}
	if result := t.db.Where("id = ?", id).Delete(&teacher); result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}
