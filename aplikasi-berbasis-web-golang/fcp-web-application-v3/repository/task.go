package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
)

type TaskRepository interface {
	Store(task *model.Task) error
	Update(taskID int, task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskRepository struct {
	filebased *filebased.Data
}

func NewTaskRepo(filebasedDb *filebased.Data) *taskRepository {
	return &taskRepository{
		filebased: filebasedDb,
	}
}

func (t *taskRepository) Store(task *model.Task) error {
	t.filebased.StoreTask(*task)

	return nil
}

func (t *taskRepository) Update(taskID int, task *model.Task) error {
		if _, err := t.GetByID(task.ID); err != nil {//mengambil task berdasarkan id dan mengcek apakah ada error
			return err
		}
		t.filebased.UpdateTask(task.ID, *task)//mengupdate data task
		return nil // TODO: replace this
}

func (t *taskRepository) Delete(id int) error {
	err := t.filebased.DeleteTask(id)// mengambil task berdasarkan id
	if err != nil {// pengecekkan error
		return err
	}
	t.filebased.DeleteTask(id)
	return nil // TODO: replace this
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	task, err := t.filebased.GetTaskByID(id)
	if err != nil { 
		return nil, err 
	}
	return task, nil // TODO: replace this
}

func (t *taskRepository) GetList() ([]model.Task, error) {
	task, err := t.filebased.GetTasks()// mengambil data task
	if err != nil {//pengecekkan error
		return nil, err
	}
	//mengembalikan data task
	return task, nil // TODO: replace this
}

func (t *taskRepository) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	taskCategories, err := t.filebased.GetTaskListByCategory(id)//mengambil data task berdasarkan id
	if err != nil {//pengecekan error
		return nil, err
	}
	//mengembalikan data task
	return taskCategories, nil // TODO: replace this
}
