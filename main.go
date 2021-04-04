package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	api "github.com/lud0v1c/go-calendar-api/api"
)

func main() {

	r := api.SetupRoutes()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	api.StartDataBase()

	r.Run(":8080")
}
