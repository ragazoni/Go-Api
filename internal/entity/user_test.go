package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Week", "jon@teste", "12345")
	assert.Nil(t, err)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "john Week", user.Name)
	assert.Equal(t, "j@j.com", user.Email)
}

func TesNewPassword(t *testing.T) {
	user, err := NewUser("John Week", "jon@teste", "12345")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.NotEqual(t, "123456", user.Password)

}
