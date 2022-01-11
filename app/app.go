package app

import (
	"github.com/avinashb98/myfin/config"
	"github.com/avinashb98/myfin/http"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.New()
)

func StartApplication() {
	router.Use(gin.Recovery())
	router.GET("/status", http.Status)

	port := config.SERVER_PORT
	err := router.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
