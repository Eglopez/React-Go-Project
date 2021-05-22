package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/LKezHn/React-Go-Project/core/models"
	"github.com/LKezHn/React-Go-Project/services"
	"github.com/gin-gonic/gin"
)

type User models.User
type UserList []User

func LoginUser(c *gin.Context) {
	user := User{}
	c.Bind(&user)
	id, password, err := services.GetUserCredentials(user.Username)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ok := services.PasswordIsValid(user.Password, password)
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Incorrect password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": services.NewToken(id),
	})
}

func GetUser(c *gin.Context) {
	id, _ := c.Get("user_id")

	user, err := services.GetUserInfo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	data, _ := json.Marshal(user)
	c.Data(http.StatusFound, "application/json", data)
}

func AddNewAccount(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	err := services.AddUser(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	data, _ := json.Marshal(user)
	c.Data(http.StatusCreated, "application/json", data)
}

func UpdateAccount(c *gin.Context) {
	user := models.User{}
	id, _ := c.Get("user_id")

	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	err = services.UpdateUser(user, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	data, _ := json.Marshal(user)
	c.Data(http.StatusAccepted, "application/json", data)
}

func DeleteAccount(c *gin.Context) {
	id, _ := c.Get("user_id")

	err := services.DeleteUser(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "account deleted",
	})
}

func ChangeAccountPassword(c *gin.Context) {
	passStruct := models.Password{}
	id, _ := c.Get("user_id")

	c.Bind(&passStruct)

	if !passStruct.IsValid() {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "New password and old password are the same",
		})
		return
	}

	encryptedPassword, err := services.GetPassword(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	if !services.PasswordIsValid(passStruct.Password, encryptedPassword) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect old password",
		})
		return
	}

	err = services.ChangePassword(passStruct.NewPassword, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "password changed",
	})

}
