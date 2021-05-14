package routers

import (
	"github.com/LKezHn/React-Go-Project/controllers"
	"github.com/gin-gonic/gin"
)

// Function to initialize the differents routes of the API
func InitRoutes(router *gin.RouterGroup) {
	router.GET("/", controllers.ShowInfo)
	router.GET("/test", controllers.Test)

	AddAccountRoutes(router)
	AddBoardRoutes(router)
}
