package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserHashPassword(t *testing.T) {
	user := User{Username: "testPerson", Email: "test@mail.com", Password: "123"}
	password := user.Password
	err := user.HashPassword()
	assert.NoError(t, err)
	err = user.Authenticate("123")
	assert.NoError(t, err)
	assert.NotEqual(t, password, user.Password)
}

func TestToLimitedUser(t *testing.T) {
	user := User{Username: "testPerson", Email: "test@mail.com", Password: "123"}
	limitedUser := user.ToLimitedUser()
	assert.Equal(t, user.ID, limitedUser.ID)
	assert.Equal(t, user.Username, limitedUser.Username)
}

func TestAuthenticate(t *testing.T) {
	tests := []struct {
		User     User
		Password string
		Err      bool
	}{
		{
			User:     User{Password: "123"},
			Password: "123",
			Err:      false,
		},
		{
			User:     User{Password: "123"},
			Password: "1234",
			Err:      true,
		},
	}
	for _, tt := range tests {
		_ = tt.User.HashPassword()
		err := tt.User.Authenticate(tt.Password)
		if tt.Err {
			assert.Error(t, err)
		}
	}
}

func TestUserValidation(t *testing.T) {
	tests := []struct {
		User  User
		Valid bool
		Err   bool
	}{
		{
			User: User{Username: "test.person1", Email: "test@mail.com", Password: "123"},
			Err:  false,
		},
		{
			User: User{Email: "test@mail.com", Password: "123"},
			Err:  true,
		},
		{
			User: User{Username: "test Person", Email: "test@mail.com", Password: "123"},
			Err:  true,
		},
		{
			User: User{Username: "testperson", Password: "123"},
			Err:  true,
		},
		{
			User: User{Username: "testperson", Email: "testmail.com", Password: "123"},
			Err:  true,
		},
		{
			User: User{Username: "testperson", Email: "test@mail.com"},
			Err:  true,
		},
		{
			User: User{Username: "testperson", Email: "test@mail.com", Password: "1 23"},
			Err:  true,
		},
		{
			User: User{Username: "test@mail.com", Email: "test@mail.com", Password: "123"},
			Err:  true,
		},
	}

	for _, tt := range tests {
		err := tt.User.Validate()
		if tt.Err {
			assert.Error(t, err)
		}
	}
}
