package services

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/LKezHn/React-Go-Project/core/models"
	"github.com/LKezHn/React-Go-Project/database"
	"github.com/go-sql-driver/mysql"
)

func AuthUser(username string) (string, string, error) {
	var (
		id       string
		password string
		e        error
	)
	db := database.Init()
	defer database.CloseConnection(db)

	err := db.QueryRow("SELECT id, str_password FROM User WHERE BINARY str_username = ?", username).Scan(&id, &password)

	if err == sql.ErrNoRows {
		e = errors.New("account doesn't exists")
	}

	return id, password, e
}

func AddUser(user *models.User) error {

	user.Password = encryptPassword(user.Password)

	db := database.Init()
	defer database.CloseConnection(db)

	query, err := db.Prepare("INSERT INTO User(str_firstName,str_lastName, str_email, str_username, str_password) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err.Error())
	}

	res, err := query.Exec(user.FirstName, user.LastName, user.Email, user.Username, user.Password)

	if driverErr, ok := err.(*mysql.MySQLError); ok {
		if driverErr.Number == 1062 {
			return errors.New("Email or username already exists")
		}
	}

	user.Id, err = res.LastInsertId()

	if err != nil {
		return errors.New("error")
	}

	return nil
}

func GetUserInfo(id interface{}) (models.User, error) {
	var (
		user models.User
		e    error
	)
	db := database.Init()
	defer database.CloseConnection(db)

	err := db.QueryRow("SELECT id, str_firstName, str_lastName, str_email, str_username FROM User WHERE User.id = ?", id).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Username,
	)
	if err == sql.ErrNoRows {
		e = errors.New("account doesn't exists")
	}

	return user, e
}
