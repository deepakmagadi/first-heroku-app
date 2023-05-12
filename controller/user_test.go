package controller

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetUser(t *testing.T) {
	var c *gin.Context
	var db *gorm.DB
	resp, _ := GetUser(c, db)
	assert.Equal(t, resp, "test user")
}
