package routers

import (
	"github.com/LKezHn/trillBackend/controllers"
	"github.com/gin-gonic/gin"
)

// Function to initialize the differents routes of the API
func InitRoutes(router *gin.RouterGroup) {
	router.GET("/", controllers.ShowInfo)
	router.GET("/test", controllers.Test)

	AddUserRoutes(router)
	AddBoardRoutes(router)
}
