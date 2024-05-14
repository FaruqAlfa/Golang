package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryAPI interface {
	AddCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
	GetCategoryByID(c *gin.Context)
	GetCategoryList(c *gin.Context)
}

type categoryAPI struct {
	categoryService service.CategoryService
}

func NewCategoryAPI(categoryRepo service.CategoryService) *categoryAPI {
	return &categoryAPI{categoryRepo}
}

func (ct *categoryAPI) AddCategory(c *gin.Context) {
	var newCategory model.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := ct.categoryService.Store(&newCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add category success"})
}

func (ct *categoryAPI) UpdateCategory(c *gin.Context) {
	var updateCategory model.Category //membuat variabel baru bertipe model category

	ctgID, err := strconv.Atoi(c.Param("id")) //membuat variabel untuk menampung id kemudian di convert ke integer
	if err != nil{// pengecekan id apakah valid atau tidak
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid Category ID"})
		return
	}

	if err := c.ShouldBindJSON(&updateCategory); err != nil{ //pengecekan apakah data yang dikirim valid atau tidak
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	updateCategory.ID = ctgID //mengupdate id dari variabel updateCategory
	err = ct.categoryService.Update(ctgID, updateCategory) // mengupdate data dari variabel updateCategory
	if err != nil {// pengecekan untuk apakah ada error
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})	
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "category update success"})//menampilkan pesan berhasil
	// TODO: answer here
}

func (ct *categoryAPI) DeleteCategory(c *gin.Context) {
	ctgID, err := strconv.Atoi(c.Param("id")) //membuat variabel untuk menampung id kemudian di convert ke integer
	if err != nil {// pengecekan id apakah ada kesalahan input atau tidak
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid category ID"})
		return
	}

	err = ct.categoryService.Delete(ctgID) //menghapus data dari variabel updateCategory
	if err != nil {// pengecekan untuk apakah ada error pada penghapusan data
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "category delete success"})//menampilkan pesan berhasil
	// TODO: answer here
}

func (ct *categoryAPI) GetCategoryByID(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid category ID"})
		return
	}

	category, err := ct.categoryService.GetByID(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (ct *categoryAPI) GetCategoryList(c *gin.Context) {
	ctgList, err := ct.categoryService.GetList()// memanggil fungsi GetList pada service category
	if err != nil { // pengecekan untuk apakah ada error atau tidak pada saat memanggil fungsi GetList
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ctgList)// menampilkan category list
	// TODO: answer here
}
