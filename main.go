package main

import (
	"os"
	"time"

	"first-heroku-app.com/controller"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

func main() {
	godotenv.Load()
	dsn := "host=localhost user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})

	r := gin.Default()
	//env := os.Getenv("ENV")
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "You have connected with develop server",
		})
	})
	r.POST("/user", controller.CreateUser)
	r.GET("/user", controller.GetUser)
	port := os.Getenv("PORT")
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080
}
