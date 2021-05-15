package services

import (
	"errors"
	"fmt"

	"github.com/LKezHn/React-Go-Project/core/models"
	"github.com/LKezHn/React-Go-Project/database"
)

type Board models.Board
type BoardList []Board

func GetBoards(id interface{}) []Board {
	var boards = BoardList{}
	db := database.Init()
	defer database.CloseConnection(db)

	rows, err := db.Query("SELECT Board.id, Board.str_name, Board.str_backgroundColor FROM Board INNER JOIN UserBoard ON UserBoard.id_board = Board.id WHERE UserBoard.id_user = ?", id)

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		board := Board{}
		rows.Scan(&board.Id, &board.Name, &board.BackgroundColor)
		boards = append(boards, board)
	}
	return boards
}

func CreateBoard(newBoard *models.Board) error {
	db := database.Init()
	defer database.CloseConnection(db)

	query, err := db.Prepare("INSERT INTO Board(str_name, str_backgroundColor) VALUES(?,?)")
	if err != nil {
		return errors.New("Firsterror")
	}

	res, err := query.Exec(newBoard.Name, newBoard.BackgroundColor)

	if err != nil {
		return errors.New("Seconderror")
	}

	newBoard.Id, err = res.LastInsertId()
	if err != nil {
		return errors.New("Thirderror")
	}

	return nil
}

func AddBoardToUser(user_id interface{}, board_id int64) error {
	db := database.Init()
	defer database.CloseConnection(db)

	query, err := db.Prepare("INSERT INTO UserBoard VALUES(?,?)")
	if err != nil {
		return errors.New("Fourtherror")
	}

	_, err = query.Exec(user_id, board_id)

	if err != nil {
		return errors.New("Fiftherror")
	}

	return nil
}
