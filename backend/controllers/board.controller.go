package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/LKezHn/React-Go-Project/core/models"
	"github.com/LKezHn/React-Go-Project/services"
	"github.com/gin-gonic/gin"
)

type Board models.Board
type BoardList []Board

func GetUserBoards(c *gin.Context) {
	id, _ := c.Get("user_id")
	boards := services.GetBoards(id)

	data, _ := json.Marshal(boards)

	c.Data(http.StatusOK, "application/json", data)
}

func AddNewBoard(c *gin.Context) {
	board := models.Board{}
	id, _ := c.Get("user_id")
	c.Bind(&board)

	err := services.CreateBoard(&board)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = services.AddBoardToUser(id, board.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Board created",
	})

}
