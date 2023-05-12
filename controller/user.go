package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

func CreateUser(c *gin.Context) {
	var dsn string
	if os.Getenv("DATABASE_URL") == "local" {
		dsn = "host=ec2-54-234-13-16.compute-1.amazonaws.com user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + ""
	} else {
		dsn = os.Getenv("DATABASE_URL")
	}
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	user := User{
		Name:      "deepak",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := db.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusCreated, user)
}

func GetUser(c *gin.Context) {
	var dsn string
	if os.Getenv("DATABASE_URL") == "local" {
		dsn = "host=ec2-54-234-13-16.compute-1.amazonaws.com user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + ""
	} else {
		dsn = os.Getenv("DATABASE_URL")
	}
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	var user []User
	err := db.Find(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, user)
}
