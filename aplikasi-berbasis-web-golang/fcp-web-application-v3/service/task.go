package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
)

type TaskService interface {
	Store(task *model.Task) error
	Update(id int, task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskService struct {
	taskRepository repo.TaskRepository
}

func NewTaskService(taskRepository repo.TaskRepository) TaskService {
	return &taskService{taskRepository}
}

func (c *taskService) Store(task *model.Task) error {
	err := c.taskRepository.Store(task)
	if err != nil {
		return err
	}

	return nil
}

func (s *taskService) Update(id int, task *model.Task) error {
	task, err := s.taskRepository.GetByID(task.ID)//mengambil task berdasarkan id
	if err != nil { //pengecekkan error
		return err 
	}
	//mengupdate task
	return s.taskRepository.Update(task.ID, task) // TODO: replace this
}

func (s *taskService) Delete(id int) error {
	task, err := s.taskRepository.GetByID(id)//mengambil task berdasarkan id
	if err != nil {//pengecekan error
		return err
	}
	//menghapus task berdasarkan id
	return s.taskRepository.Delete(task.ID) // TODO: replace this
}

func (s *taskService) GetByID(id int) (*model.Task, error) {
	task, err := s.taskRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *taskService) GetList() ([]model.Task, error) {
	task, err := s.taskRepository.GetList()//mengambil data task 
	if err != nil {//pengecekan error
		return nil, err
	} 
	//mengembailkan data list task
	return task, nil // TODO: replace this
}

func (s *taskService) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	task, err := s.taskRepository.GetTaskCategory(id)//mengambil data task berdasarkan id
	if err != nil {//pengecekan error
		return nil, err
	} // TODO: replace this
	return task, nil//mengembalikan nilai task berdasarkan id
}
