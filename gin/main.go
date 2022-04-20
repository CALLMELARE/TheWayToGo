package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	username string
	password string
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello gin!",
	})
}

func main() {
	r := gin.Default()
	r.GET("/hello", hello)
	r.GET("/user", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBind(u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		}
	})
	r.Run()
}
