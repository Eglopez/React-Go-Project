package controllers

import (
	"net/http"

	"github.com/LKezHn/trillBackend/core"
	"github.com/gin-gonic/gin"
)

// Show the API's endpoints info
func ShowInfo(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Trill API v1.0",
		},
	)
}

func Test(c *gin.Context) {
	code := core.GenerateCode()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
	})
}
