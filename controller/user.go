package controller

import (
	"fmt"
	"github.com/avinashb98/myfin/service/user"
	"github.com/gin-gonic/gin"
)

type UserInput struct {
	Handle   string `json:"handle"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserController interface {
	CreateUser(c *gin.Context) error
	GetUserByHandle(c *gin.Context) (*user.User, error)
}

type userController struct {
	userService user.Service
}

func NewUserController(userService user.Service) UserController {
	return &userController{
		userService: userService,
	}
}

func (u *userController) CreateUser(c *gin.Context) error {
	var input UserInput
	err := c.ShouldBind(&input)
	if err != nil {
		return fmt.Errorf("invalid user details")
	}
	_user := user.User{
		Handle: input.Handle,
		Name:   input.Name,
		Email:  input.Email,
	}
	return u.userService.CreateUser(c, _user, input.Password)
}

func (u *userController) GetUserByHandle(c *gin.Context) (*user.User, error) {
	var input UserInput
	err := c.ShouldBind(&input)
	if err != nil {
		return nil, fmt.Errorf("invalid user handle")
	}
	return u.userService.GetUserByHandle(c, input.Handle)
}
