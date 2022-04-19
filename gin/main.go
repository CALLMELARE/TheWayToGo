package main

import "github.com/gin-gonic/gin"

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello gin!",
	})
}

func main() {
	r := gin.Default()
	r.GET("/hello", hello)
	r.Run()
}
