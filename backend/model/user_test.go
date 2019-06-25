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
	assert.NotEqual(t, password, user.Password)
}

func TestUserValidation(t *testing.T) {
	tests := []struct {
		User  User
		Valid bool
		Err   bool
	}{
		{
			User:  User{Username: "testPerson", Email: "test@mail.com", Password: "123"},
			Valid: true,
			Err:   false,
		},
		{
			User:  User{Email: "test@mail.com", Password: "123"},
			Valid: false,
			Err:   true,
		},
		{
			User:  User{Username: "test Person", Email: "test@mail.com", Password: "123"},
			Valid: false,
			Err:   true,
		},
		{
			User:  User{Username: "testPerson", Password: "123"},
			Valid: false,
			Err:   true,
		},
		{
			User:  User{Username: "testPerson", Email: "testmail.com", Password: "123"},
			Valid: false,
			Err:   true,
		},
		{
			User:  User{Username: "testPerson", Email: "test@mail.com"},
			Valid: false,
			Err:   true,
		},
		{
			User:  User{Username: "testPerson", Email: "test@mail.com", Password: "1 23"},
			Valid: false,
			Err:   true,
		},
		{
			User: User{Username: "test@mail.com", Email: "test@mail.com", Password: "123"},
			Valid: false,
			Err: true,
		},
	}

	for _, tt := range tests {
		valid, err := tt.User.Validate()
		assert.Equal(t, tt.Valid, valid)
		if tt.Err {
			assert.Error(t, err)
		}
	}
}
