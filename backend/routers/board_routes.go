package routers

import (
	"github.com/LKezHn/React-Go-Project/controllers"
	"github.com/LKezHn/React-Go-Project/middlewares"
	"github.com/gin-gonic/gin"
)

func AddBoardRoutes(router *gin.RouterGroup) {

	router.GET("/boards", middlewares.TokenAuthMiddleware(), controllers.GetUserBoards)
	router.POST("/boards", middlewares.TokenAuthMiddleware(), controllers.AddNewBoard)
}
