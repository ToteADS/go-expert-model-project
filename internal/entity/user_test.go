package entity

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	user, err := NewUser("Tote Araujo", "tote@tapp.dev.br", "password")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Tote Araujo", user.Name)
	assert.Equal(t, "tote@tapp.dev.br", user.Email)
	
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Tote Araujo", "tote@tapp.dev.br", "password")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.True(t, user.ValidatePassword("password"))
	assert.False(t, user.ValidatePassword("123456"))
	assert.NotEqual(t, "password", user.Password)
}