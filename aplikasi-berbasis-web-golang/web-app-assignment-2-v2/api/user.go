package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserAPI interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetUserTaskCategory(c *gin.Context)
}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userService service.UserService) *userAPI {
	return &userAPI{userService}
}

func (u *userAPI) Register(c *gin.Context) {
	var user model.UserRegister

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	if user.Email == "" || user.Password == "" || user.Fullname == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("register data is empty"))
		return
	}

	var recordUser = model.User{
		Fullname: user.Fullname,
		Email:    user.Email,
		Password: user.Password,
	}

	recordUser, err := u.userService.Register(&recordUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, model.NewSuccessResponse("register success"))
}

func (u *userAPI) Login(c *gin.Context) {
	var inputUser model.UserLogin
	err := c.ShouldBindJSON(&inputUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	if inputUser.Email == "" || inputUser.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.NewErrorResponse("email or password is empty"))
		return
	}

	var user = model.User{
		Email:    inputUser.Email,
		Password: inputUser.Password,
	}

	token, err := u.userService.Login(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	c.SetCookie("session_token", *token, expirationTime.Second(), "/", "", false, true)
	
	id, _ := c.Get("id")
	
	c.JSON(http.StatusOK, gin.H{
		"user_id":id,
		"message":"login success",
	})
// TODO: answer here
}

func (u *userAPI) GetUserTaskCategory(c *gin.Context) {
	userTaskCategories, err := u.userService.GetUserTaskCategory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		return
	}

	// Kirim respons dengan data yang diterima
	c.JSON(http.StatusOK, userTaskCategories)
	// TODO: answer here
}
