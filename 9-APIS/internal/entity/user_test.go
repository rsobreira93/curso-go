package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "john@email.com", "123456")

	assert.Nil(t, err)

	assert.NotEmpty(t, user)

	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)

	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john@email.com", user.Email)

}

func Test_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "jhon@email.com", "123456")

	assert.Nil(t, err)

	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))

	assert.NotEqual(t, "123456", user.Password)
}
