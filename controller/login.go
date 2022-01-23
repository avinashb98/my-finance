package controller

import (
	"fmt"
	"github.com/avinashb98/myfin/service/auth"
	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Handle   string `json:"handle"`
	Password string `json:"password"`
}

type LoginController interface {
	Login(c *gin.Context) (string, error)
}

type loginController struct {
	authService auth.Service
}

func LoginHandler(authService auth.Service) LoginController {
	return &loginController{
		authService: authService,
	}
}

func (ctrl *loginController) Login(c *gin.Context) (string, error) {
	var credential Credentials
	err := c.ShouldBind(&credential)
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}
	isUserAuthenticated, err := ctrl.authService.IsAuthenticated(c, credential.Handle, credential.Password)
	if err != nil {
		return "", err
	}
	if !isUserAuthenticated {
		return "", fmt.Errorf("user unauthorised, incorrect handle or password")
	}
	return ctrl.authService.GenerateToken(credential.Handle, true)
}
