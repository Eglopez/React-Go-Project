package routers

import (
	"github.com/LKezHn/React-Go-Project/controllers"
	"github.com/gin-gonic/gin"
)

func AddAccountRoutes(router *gin.RouterGroup) {
	router.POST("/account/signup", controllers.AddNewAccount)
	router.POST("/account/login", controllers.LoginUser)
	router.GET("/account", controllers.GetUser)
	router.PUT("users/:id", controllers.UpdateUser)
	router.DELETE("users/:id", controllers.DeleteUser)
}
