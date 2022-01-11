package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type StatusDTO struct {
	Message string `json:"message"`
}

func Status(c *gin.Context) {
	c.JSON(http.StatusOK, StatusDTO{
		Message: "Healthy!",
	})
}
