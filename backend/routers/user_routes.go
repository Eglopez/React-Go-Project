package routers

import (
	"github.com/LKezHn/React-Go-Project/controllers"
	"github.com/gin-gonic/gin"
)

func AddUserRoutes(router *gin.RouterGroup) {

	router.GET("/users", controllers.GetAllUsers)
	router.POST("/users", controllers.AddNewUser)
	router.GET("/users/:id", controllers.GetUser)
	router.PUT("users/:id", controllers.UpdateUser)
	router.DELETE("users/:id", controllers.DeleteUser)
}
