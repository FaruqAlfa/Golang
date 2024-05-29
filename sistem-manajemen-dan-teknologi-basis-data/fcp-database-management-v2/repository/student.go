package repository

import (
	"a21hc3NpZ25tZW50/model"
	// "fmt"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
	FetchWithClass() (*[]model.StudentClass, error)
}

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	//mengambil semua data student dari database kemudian di scan kedalam slice student
	var students []model.Student
	if err := s.db.Find(&students).Error; err != nil {
		return nil, err
	}

	return students, nil
	// return []model.Student{}, nil // TODO: replace this
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	//menyimpan data student kedalam database
	if err := s.db.Create(&student).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (s *studentRepoImpl) Update(id int, student *model.Student) error {
	//mengupdate data student berdasarkan id
	if err := s.db.Model(&student).Where("id = ?", id).Updates(student).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (s *studentRepoImpl) Delete(id int) error {
	//menghapus data student berdasarkan id
	if err := s.db.Where("id = ?", id).Delete(&model.Student{}).Error; err != nil {
		return err
	}

	return nil
	// return fmt.Errorf("not implement") // TODO: replace this
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	//mengambil data student berdasarkan id dari database
	var student model.Student
	err := s.db.First(&student, id).Error
	if err != nil {
		return nil, err
	}

	return &student, nil
	// return nil, nil // TODO: replace this
}

func (s *studentRepoImpl) FetchWithClass() (*[]model.StudentClass, error) {
	//mengambil semua data student beserta class-nya
	var studentClasses []model.StudentClass
	err := s.db.Table("students").
		Select("students.name, students.address, classes.name as class_name, classes.professor, classes.room_number").
		Joins("left join classes on students.class_id = classes.id").
		Scan(&studentClasses).Error
	if err != nil {
		return nil, err
	}

	if len(studentClasses) == 0 {
		return &[]model.StudentClass{}, nil
	}

	return &studentClasses, nil// TODO: replace this
}
