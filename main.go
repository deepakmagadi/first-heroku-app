package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := gin.Default()
	env := os.Getenv("ENV")
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "You have connected with " + env + " server",

		})
	})
	port := os.Getenv("PORT")
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080
}
