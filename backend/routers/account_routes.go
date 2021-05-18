package routers

import (
	"github.com/LKezHn/React-Go-Project/controllers"
	"github.com/LKezHn/React-Go-Project/middlewares"
	"github.com/gin-gonic/gin"
)

func AddAccountRoutes(router *gin.RouterGroup) {
	router.POST("/account/signup", controllers.AddNewAccount)
	router.POST("/account/login", controllers.LoginUser)
	router.GET("/account", middlewares.TokenAuthMiddleware(), controllers.GetUser)
	router.PUT("/account", middlewares.TokenAuthMiddleware(), controllers.UpdateAccount)
}
