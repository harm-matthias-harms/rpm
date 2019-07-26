package storage

import (
	"context"
	"testing"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	resetUserDatabase()
	// Creates a User
	user := &model.User{Username: "test", Email: "test@mail.com", Password: "test"}
	err := CreateUser(user)
	assert.NoError(t, err)
	// Can't create twice
	err = CreateUser(user)
	assert.Error(t, err)
}

func TestFindUser(t *testing.T) {
	resetUserDatabase()
	// Find user by id
	// First create him because no user is inserted
	user := &model.User{Username: "test", Email: "test@mail.com"}
	// Find user by username
	userWithUsername := &model.User{Username: "test"}
	// Find user by email for login
	userWithEmail := &model.User{Username: "test@mail.com"}
	err := CreateUser(user)
	assert.NoError(t, err)
	userFound, err := FindUser(user)
	assert.NoError(t, err)
	//set nil to find only by id
	userFound.Username = ""
	userFound.Email = ""
	tests := []*model.User{userFound, userWithUsername, userWithEmail}
	for _, tt := range tests {
		userFound, err = FindUser(tt)
		assert.NoError(t, err)
		assert.Equal(t, "test", userFound.Username)
		assert.Equal(t, "test@mail.com", userFound.Email)
	}
	// Dont't find not existing user
	userNotExist := &model.User{ Username: "test1"}
	_, err = FindUser(userNotExist)
	assert.Error(t, err)
}

func resetUserDatabase() {
	_ = SetMongoDatabase()
	_ = userCollection().Drop(context.Background())
	CreateUserIndexes()
}
