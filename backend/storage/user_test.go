package storage

import (
	"testing"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateUser(t *testing.T) {
	// Creates a User
	user := &model.User{Username: "test", Email: "test1@mail.com", Password: "test"}
	resetUserDatabase(user)
	err := CreateUser(nil, user)
	assert.NoError(t, err)
	// Can't create twice
	err = CreateUser(nil, user)
	assert.Error(t, err)
}

func TestFindUser(t *testing.T) {
	// Find user by id
	// First create him because no user is inserted
	user := &model.User{Username: "test", Email: "test1@mail.com"}
	resetUserDatabase(user)
	// Find user by username
	userWithUsername := &model.User{Username: "test"}
	// Find user by email for login
	userWithEmail := &model.User{Username: "test1@mail.com"}
	err := CreateUser(nil, user)
	assert.NoError(t, err)
	userFound, err := FindUser(nil, user)
	assert.NoError(t, err)
	//set nil to find only by id
	userFound.Username = ""
	userFound.Email = ""
	tests := []*model.User{userFound, userWithUsername, userWithEmail}
	for _, tt := range tests {
		userFound, err = FindUser(nil, tt)
		assert.NoError(t, err)
		assert.Equal(t, "test", userFound.Username)
		assert.Equal(t, "test1@mail.com", userFound.Email)
	}
	// Dont't find not existing user
	userNotExist := &model.User{Username: "test1"}
	_, err = FindUser(nil, userNotExist)
	assert.Error(t, err)
}

func resetUserDatabase(user *model.User) {
	_ = SetMongoDatabase()
	_, _ = userCollection().DeleteOne(nil, bson.M{"username": user.Username})
}
