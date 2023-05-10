package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	resp := GetUser()
	assert.Equal(t, resp, "test user")
}
