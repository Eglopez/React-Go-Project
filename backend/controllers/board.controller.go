package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/LKezHn/trillBackend/core/models"
	"github.com/gin-gonic/gin"
)

type Board models.Board
type BoardList []Board

var list = BoardList{
	Board{1, "Base de Datos 1", "#ff00ff"},
	Board{2, "Inteligencia Artificial", "#f6f6f6"},
	Board{3, "Google", "#090909"},
}

func GetAllBoard(c *gin.Context) {
	data, _ := json.Marshal(list)
	c.Data(http.StatusOK, "application/json", data)
}

func AddNewBoard(c *gin.Context) {
	var board Board
	c.Bind(&board)
	list = append(list, board)
	c.JSON(http.StatusCreated, board)
}
