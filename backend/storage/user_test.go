package storage

import (
	"testing"

	"github.com/harm-matthias-harms/rpm/backend/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestUser(t *testing.T) {
	// setup
	_ = SetMongoDatabase()

	// model
	user := &model.User{Username: "test", Email: "test1@mail.com", Password: "test", Code: "testCode"}

	// tests
	t.Run("create user", func(t *testing.T) {
		err := CreateUser(nil, user)
		assert.NoError(t, err)
		// Can't create twice
		err = CreateUser(nil, user)
		assert.Error(t, err)
	})

	t.Run("find user", func(t *testing.T) {
		// submodel setup
		userWithUsername := &model.User{Username: "test"}
		userWithEmail := &model.User{Username: "test1@mail.com"}
		userWithCode := &model.User{Code: "testCode"}
		userNotExist := &model.User{Username: "test1"}

		// find base user
		userFound, err := FindUser(nil, user)

		//set nil to find only by id
		userFound.Username = ""
		userFound.Email = ""

		// find different user
		tests := []*model.User{userFound, userWithUsername, userWithEmail, userWithCode}
		for _, tt := range tests {
			userFound, err = FindUser(nil, tt)
			assert.NoError(t, err)
			assert.Equal(t, "test", userFound.Username)
			assert.Equal(t, "test1@mail.com", userFound.Email)
		}

		// Dont't find not existing user
		_, err = FindUser(nil, userNotExist)
		assert.Error(t, err)
	})

	// cleanup
	resetUserDatabase(user)
}

func resetUserDatabase(user *model.User) {
	_, _ = userCollection().DeleteMany(nil, bson.M{"username": user.Username})
}
