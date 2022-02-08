package app

import (
	"context"
	"github.com/avinashb98/myfin/config"
	"github.com/avinashb98/myfin/controller"
	"github.com/avinashb98/myfin/datasources/mongo"
	"github.com/avinashb98/myfin/middleware"
	userRepo "github.com/avinashb98/myfin/repository/user"
	"github.com/avinashb98/myfin/service/auth"
	userService "github.com/avinashb98/myfin/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	router = gin.New()
)

func StartApplication() {
	ctx := context.Background()
	conf := config.GetConfig()
	router.Use(gin.Recovery())
	router.GET("/status", controller.Status)

	mongoClient, err := mongo.NewClient(&conf.Mongo)
	if err != nil {
		panic(err)
	}
	err = mongoClient.Connect()
	if err != nil {
		panic(err)
	}
	mongoDB := mongo.NewDatabase(&conf.Mongo, mongoClient)

	userR := userRepo.NewRepository(ctx, mongoDB)
	userS := userService.NewService(userR)
	userController := controller.NewUserController(userS)

	authService := auth.NewService(conf.JWT, userR)
	loginHandler := controller.LoginHandler(authService)
	authMiddleware := middleware.AuthorizeJWT(authService)

	apiV1Router := router.Group("/api/v1")
	{
		apiV1Router.POST("/login", func(context *gin.Context) {
			token, err := loginHandler.Login(context)
			if err != nil {
				context.JSON(http.StatusUnauthorized, gin.H{
					"message": err.Error(),
				})
				return
			}
			context.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		})

		apiV1Router.GET("/user/:handle", authMiddleware, func(c *gin.Context) {
			_user, err := userController.GetUserByHandle(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, _user)
		})

		apiV1Router.POST("/user", func(c *gin.Context) {
			err := userController.CreateUser(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			}
			c.JSON(http.StatusCreated, gin.H{
				"message": "User successfully created",
			})
		})
	}

	port := config.ServerPort
	err = router.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
