package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LKezHn/React-Go-Project/core/models"

	"github.com/LKezHn/React-Go-Project/database"
	enc "github.com/LKezHn/React-Go-Project/libs/EncryptManager"
	em "github.com/LKezHn/React-Go-Project/libs/ErrorManager"
	tm "github.com/LKezHn/React-Go-Project/libs/TokenManager"
	"github.com/gin-gonic/gin"
)

type User models.User
type UserList []User

func LoginUser(c *gin.Context) {
	db := database.Init()
	defer database.CloseConnection(db)

	user := User{}
	c.Bind(&user)
	id, password := "", ""

	err := db.QueryRow("SELECT id, str_password FROM User WHERE BINARY str_username = ?", user.Username).Scan(&id, &password)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Account not found",
		})
		return
	}

	if enc.CheckPassword(user.Password, password) {
		c.JSON(http.StatusFound, gin.H{
			"token": tm.NewToken(id),
		})
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Incorrect password",
		})
	}

}

func GetUser(c *gin.Context) {
	db := database.Init()
	defer database.CloseConnection(db)
	var user User
	token := c.GetHeader("Authorization")
	fmt.Println(token)
	id, err := tm.ValidateToken(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	row, err := db.Query("SELECT id, str_firstName, str_lastName, str_email, str_username FROM User WHERE User.id = ?", id)
	em.ErrorCheck(err)

	for row.Next() {
		err = row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Username)
		em.ErrorCheck(err)
	}

	data, _ := json.Marshal(user)
	c.Data(http.StatusFound, "application/json", data)

}

func AddNewAccount(c *gin.Context) {
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
