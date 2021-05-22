package services

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/LKezHn/React-Go-Project/core/models"
	"github.com/LKezHn/React-Go-Project/database"
	"github.com/go-sql-driver/mysql"
)

func GetUserCredentials(username string) (string, string, error) {
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

func GetPassword(id interface{}) (string, error) {
	var password string
	db := database.Init()
	defer database.CloseConnection(db)

	err := db.QueryRow("SELECT str_password FROM User WHERE User.id = ?", id).Scan(&password)

	if err == sql.ErrNoRows {
		return "", errors.New("User id doesn't exists")
	}

	return password, nil

}

func AddUser(user *models.User) error {

	user.Password = encryptPassword(user.Password)

	db := database.Init()
	defer database.CloseConnection(db)

	query, err := db.Prepare("INSERT INTO User(str_firstName,str_lastName, str_email, str_username, str_password) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return errors.New("error")
	}

	res, err := query.Exec(user.FirstName, user.LastName, user.Email, user.Username, user.Password)

	if driverErr, ok := err.(*mysql.MySQLError); ok {
		if driverErr.Number == 1062 {
			if strings.Contains(driverErr.Error(), "str_email") {
				return errors.New("Email already exists")
			}

			if strings.Contains(driverErr.Error(), "str_username") {
				return errors.New("Username already exists")
			}
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

func UpdateUser(user models.User, id interface{}) error {
	db := database.Init()
	defer database.CloseConnection(db)

	query, err := db.Prepare(`
		UPDATE User SET str_firstName = ?, str_lastName = ?, str_email = ?, str_username = ? WHERE User.id = ?`)

	if err != nil {
		return errors.New("Error in query")
	}

	res, err := query.Exec(user.FirstName, user.LastName, user.Email, user.Username, id)

	if driverErr, ok := err.(*mysql.MySQLError); ok {
		if driverErr.Number == 1062 {
			if strings.Contains(driverErr.Error(), "str_email") {
				return errors.New("Email already exists")
			}

			if strings.Contains(driverErr.Error(), "str_username") {
				return errors.New("Username already exists")
			}
		}
	}

	_, err = res.LastInsertId()

	if err != nil {
		return errors.New("error")
	}

	return nil
}

func DeleteUser(id interface{}) error {
	db := database.Init()
	defer database.CloseConnection(db)

	query, err := db.Prepare(`
		DELETE FROM User WHERE User.id = ?
	`)

	if err != nil {
		return errors.New("Error in query")
	}

	res, err := query.Exec(id)

	if err != nil {
		return errors.New("error")
	}
	_, err = res.LastInsertId()

	if err != nil {
		return errors.New("error")
	}

	return nil
}

func ChangePassword(newPass string, id interface{}) error {
	db := database.Init()
	defer database.CloseConnection(db)

	newPass = encryptPassword(newPass)

	query, err := db.Prepare(`
		UPDATE User SET str_password = ? WHERE User.id = ?
	`)

	if err != nil {
		return errors.New("Error in query")
	}

	res, err := query.Exec(newPass, id)

	if err != nil {
		return errors.New("error")
	}
	_, err = res.LastInsertId()

	if err != nil {
		return errors.New("error")
	}

	return nil
}
