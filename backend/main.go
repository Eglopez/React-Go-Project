package main

import (
	"net/http"

	"github.com/LKezHn/React-Go-Project/backend/routers"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/api/v1")
	})

	v1 := r.Group("/api/v1")
	routers.InitRoutes(v1)

	r.Run(":4400")
}
