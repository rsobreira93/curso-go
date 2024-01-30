package database

import (
	"testing"

	"github.com/rsobreira93/curso-go/9-APIS/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("John", "j@j.com", "123456")

	userDB := NewUser(db)

	err = userDB.Create(user)

	assert.Nil(t, err)

	var userFound entity.User

	err = db.First(&userFound, "id = $1", user.ID).Error

	assert.Nil(t, err)

	assert.Equal(t, user.ID, userFound.ID)
}
