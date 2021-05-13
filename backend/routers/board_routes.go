package routers

import (
	"github.com/LKezHn/React-Go-Project/controllers"
	"github.com/gin-gonic/gin"
)

func AddBoardRoutes(router *gin.RouterGroup) {

	router.GET("/boards", controllers.GetAllBoard)
	router.POST("/boards", controllers.AddNewBoard)
}
