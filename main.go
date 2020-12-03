package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/users", getUsersHandler)
	r.Run()
}

func getUsersHandler(c *gin.Context){

	c.JSON(200, gin.H{
		"status": "ok",
	})

}