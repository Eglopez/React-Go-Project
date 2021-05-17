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

	rows, err := db.Query(`
		SELECT Board.id, Board.str_name, Board.str_backgroundColor 
		FROM Board INNER JOIN UserBoard ON UserBoard.id_board = Board.id
		INNER JOIN BoardMember ON BoardMember.id_board = Board.id 
		WHERE UserBoard.id_user = ? OR BoardMember.id_member = ?`, id, id)

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

func GetMembers(id string) []models.User {
	members := []models.User{}

	db := database.Init()
	defer database.CloseConnection(db)

	rows, err := db.Query(`
	SELECT User.id, User.str_firstName, str_lastName, str_email, str_username FROM User 
	INNER JOIN BoardMember ON BoardMember.id_member = User.id WHERE BoardMember.id_board = ? `, id)

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		member := models.User{}
		rows.Scan(&member.Id, &member.FirstName, &member.LastName, &member.Email, &member.Username)
		members = append(members, member)
	}
	return members

}

func CreateBoard(newBoard *models.Board) error {
	db := database.Init()
	defer database.CloseConnection(db)

	query, err := db.Prepare("INSERT INTO Board(str_name, str_backgroundColor) VALUES(?,?)")
	if err != nil {
		return errors.New("Error in query declaration")
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
		return errors.New("Error in query declaration")
	}

	_, err = query.Exec(user_id, board_id)

	if err != nil {
		return errors.New("Fiftherror")
	}

	return nil
}

func AddMemberToBoard(member_id interface{}, board_id int64) error {
	db := database.Init()
	defer database.CloseConnection(db)

	query, err := db.Prepare("INSERT INTO BoardMember VALUES (?,?)")
	if err != nil {
		return errors.New("Error in query declaration")
	}

	_, err = query.Exec(board_id, member_id)

	if err != nil {
		return errors.New("Error in insertion")
	}

	return nil
}
