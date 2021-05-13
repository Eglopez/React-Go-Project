package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/LKezHn/trillBackend/core/models"
	"github.com/LKezHn/trillBackend/database"
	enc "github.com/LKezHn/trillBackend/libs/EncryptManager"
	em "github.com/LKezHn/trillBackend/libs/ErrorManager"
	"github.com/gin-gonic/gin"
)

type User models.User
type UserList []User

func GetAllUsers(c *gin.Context) {
	var users = UserList{}
	db := database.Init()
	defer database.CloseConnection(db)
	rows, err := db.Query("SELECT id, str_firstName, str_lastName, str_email, str_username FROM User")
	em.ErrorCheck(err)

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Username)
		em.ErrorCheck(err)
		users = append(users, user)
	}

	data, _ := json.Marshal(users)
	c.Data(http.StatusOK, "application/json", data)
}

func GetUser(c *gin.Context) {
	db := database.Init()
	defer database.CloseConnection(db)
	var user User
	id := c.Param("id")

	row, err := db.Query("SELECT id, str_firstName, str_lastName, str_email, str_username FROM User WHERE User.id = ?", id)
	em.ErrorCheck(err)

	for row.Next() {
		err = row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Username)
		em.ErrorCheck(err)
	}

	data, _ := json.Marshal(user)
	c.Data(http.StatusFound, "application/json", data)

}

func AddNewUser(c *gin.Context) {
	var user User
	c.Bind(&user)
	user.Password = enc.EncryptPassword(user.Password)
	db := database.Init()
	defer database.CloseConnection(db)

	query, err := db.Prepare("INSERT INTO User(str_firstName,str_lastName, str_email, str_username, str_password) VALUES (?, ?, ?, ?, ?)")
	em.ErrorCheck(err)

	res, err := query.Exec(user.FirstName, user.LastName, user.Email, user.Username, user.Password)
	em.ErrorCheck(err)

	user.Id, err = res.LastInsertId()

	em.ErrorCheck(err)
	data, _ := json.Marshal(user)
	c.Data(http.StatusCreated, "application/json", data)
}

func UpdateUser(c *gin.Context) {
	var user User
	id := c.Param("id")
	c.Bind(&user)
	db := database.Init()
	defer database.CloseConnection(db)

	query, err := db.Prepare("UPDATE User SET str_firstName=?, str_lastName=?, str_email=?, str_username=? WHERE id=?")
	em.ErrorCheck(err)

	res, err := query.Exec(user.FirstName, user.LastName, user.Email, user.Username, id)
	em.ErrorCheck(err)

	user.Id, err = res.LastInsertId()
	em.ErrorCheck(err)

	data, _ := json.Marshal(user)

	c.Data(http.StatusOK, "application/json", data)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	db := database.Init()
	defer database.CloseConnection(db)

	query, err := db.Prepare("DELETE FROM User WHERE id=?")
	em.ErrorCheck(err)

	query.Exec(id)

	c.JSON(http.StatusOK, gin.H{
		"status": "Deleted",
	})
}
