package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
)

type CategoryService interface {
	Store(category *model.Category) error
	Update(id int, category model.Category) error
	Delete(id int) error
	GetByID(id int) (*model.Category, error)
	GetList() ([]model.Category, error)
}

type categoryService struct {
	categoryRepository repo.CategoryRepository
}

func NewCategoryService(categoryRepository repo.CategoryRepository) CategoryService {
	return &categoryService{categoryRepository}
}

func (c *categoryService) Store(category *model.Category) error {
	err := c.categoryRepository.Store(category)
	if err != nil {
		return err
	}

	return nil
}

func (c *categoryService) Update(id int, category model.Category) error {
	if _, err := c.categoryRepository.GetByID(id); err != nil {//mengambil data berdasarkan id dan pengecekan error
		return err
	}
	c.categoryRepository.Update(id, category)//mengupdate data category
	return nil
	 // TODO: replace this
}

func (c *categoryService) Delete(id int) error {
	err := c.categoryRepository.Delete(id)
	if err != nil{
		return err
	}
	return nil // TODO: replace this
}

func (c *categoryService) GetByID(id int) (*model.Category, error) {
	category, err := c.categoryRepository.GetByID(id)//mengambil data berdasarkan id
	if err != nil {//pengecekan error
		return nil, err
	}

	return category, nil//mengembalikan data berdasarkan id
}

func (c *categoryService) GetList() ([]model.Category, error) {
	category, err := c.categoryRepository.GetList()//mengambil data category
	if err != nil { //pengecekan error
		return nil, err 
	} // TODO: replace this
	return category, nil//mengembalikan data list category
}
