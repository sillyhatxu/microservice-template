package db

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

var once sync.Once

func setup() {
	InitialDatabase()
}

func TestInsert(t *testing.T) {
	once.Do(setup)
	user := User{
		LoginName:        "LoginName",
		Password:         "Password",
		UserName:         "UserName",
		Status:           false,
		Platform:         "Platform",
		CreatedTime:      time.Time{},
		LastModifiedTime: time.Time{},
	}
	err := Insert(user)
	assert.Nil(t, err)
}

func TestUpdate(t *testing.T) {
	once.Do(setup)
	user := User{
		Id:               1,
		LoginName:        "LoginNameUpdate",
		Password:         "PasswordUpdate",
		UserName:         "UserNameUpdate",
		Status:           true,
		Platform:         "PlatformUpdate",
		CreatedTime:      time.Time{},
		LastModifiedTime: time.Time{},
	}
	err := Update(user)
	assert.Nil(t, err)
}

func TestFindById(t *testing.T) {
	once.Do(setup)
	user := User{
		Id:               1,
		LoginName:        "LoginNameUpdate",
		Password:         "PasswordUpdate",
		UserName:         "UserNameUpdate",
		Status:           true,
		Platform:         "PlatformUp",
		CreatedTime:      time.Time{},
		LastModifiedTime: time.Time{},
	}
	u, err := FindById(int64(1))
	assert.Nil(t, err)
	assert.EqualValues(t, user.Id, u.Id)
	assert.EqualValues(t, user.LoginName, u.LoginName)
	assert.EqualValues(t, user.Password, u.Password)
	assert.EqualValues(t, user.UserName, u.UserName)
	assert.EqualValues(t, user.Status, u.Status)
	assert.EqualValues(t, user.Platform, u.Platform)
}
