package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
	// "fmt"
)

type CategoryRepository interface {
	Store(Category *model.Category) error
	Update(id int, category model.Category) error
	Delete(id int) error
	GetByID(id int) (*model.Category, error)
	GetList() ([]model.Category, error)
}

type categoryRepository struct {
	filebasedDb *filebased.Data
}

func NewCategoryRepo(filebasedDb *filebased.Data) *categoryRepository {
	return &categoryRepository{filebasedDb}
}

func (c *categoryRepository) Store(Category *model.Category) error {
	c.filebasedDb.StoreCategory(*Category)
	return nil
}

func (c *categoryRepository) Update(id int, category model.Category) error {
	if _, err := c.GetByID(id); err != nil {//mengambil data berdasarkan id
		return err
	}
	category.ID = id 
	c.filebasedDb.UpdateCategory(id, category)//mengupdate data category
	return nil // TODO: replace this
}

func (c *categoryRepository) Delete(id int) error {
	// return fmt.Errorf("not implement") 
	if _, err := c.GetByID(id); err != nil {//mengambil data berdasarkan id dan pengecekan error
		return err
	}
	c.filebasedDb.DeleteCategory(id)//menghapus data category berdasarkan id
	return nil// TODO: replace this
}

func (c *categoryRepository) GetByID(id int) (*model.Category, error) {
	category, err := c.filebasedDb.GetCategoryByID(id)

	return category, err
}

func (c *categoryRepository) GetList() ([]model.Category, error) {
	category, err := c.filebasedDb.GetCategories()//mengambil data category
	if err != nil {//pengecaekan error
		return nil, err
	}
	//mengembali data list category 
	return category, nil // TODO: replace this
}
