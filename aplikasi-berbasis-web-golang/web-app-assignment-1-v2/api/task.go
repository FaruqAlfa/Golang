package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskAPI interface {
	AddTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	GetTaskByID(c *gin.Context)
	GetTaskList(c *gin.Context)
	GetTaskListByCategory(c *gin.Context)
}

type taskAPI struct {
	taskService service.TaskService
}

func NewTaskAPI(taskRepo service.TaskService) *taskAPI {
	return &taskAPI{taskRepo}
}

func (t *taskAPI) AddTask(c *gin.Context) {
	var newTask model.Task 
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return	
	}

	err := t.taskService.Store(&newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add task success"})
}

func (t *taskAPI) UpdateTask(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))//membuat variabel untuk menampung id kemudian di convert ke integer
	if err != nil {// pengecekan id apakah valid atau tidak
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	var updatedTask model.Task //membuuat variabel baru bertia tipe model task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {//pengecekan apakah data yang dikirim valid atau tidak
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	updatedTask.ID = taskID //mengupdate id dari variabel updatedTask
	err = t.taskService.Update(&updatedTask)//mengupdate data dari variabel updatedTask
	if err != nil {//pengecekan apakah terdapat error atau tidak
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "update task success"})//menampilkan pesan berhasil
	// TODO: answer here
}

func (t *taskAPI) DeleteTask(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))// membuat variabel untuk menampung id kemudian di convert ke integer
	if err != nil {// pengecekan id apakah valid atau tidak
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	err = t.taskService.Delete(taskID) //menghapus data dari variabel updatedTask
	if err != nil {//pengcekan apakah terdapat error atau tidak 
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	//menampilkan pesan task berhasil dihapus
	c.JSON(http.StatusOK, model.SuccessResponse{Message: "delete task success"})
	// TODO: answer here
}

func (t *taskAPI) GetTaskByID(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	task, err := t.taskService.GetByID(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (t *taskAPI) GetTaskList(c *gin.Context) {
	tasks, err := t.taskService.GetList()// memanggil fungsi GetList pada service task
	if err != nil {// pengecekan apakah terdapat error atau tidak
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)//menampilkan task list
	// TODO: answer here
}

func (t *taskAPI) GetTaskListByCategory(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("id"))//membuat variabel untuk menampung id kemudian di convert ke integer
	if err != nil {// pengecekan id apakah valid atau tidak
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid category ID"})
		return
	}

	tasks, err := t.taskService.GetTaskCategory(categoryID) //memanggil fungsi GetTaskCategory pada service task
	if err != nil {// pengecekan apakah terdapat error atau tidak
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)//menampilkan task list berdasarkan category
	// TODO: answer here
}
