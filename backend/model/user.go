package model

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is the type of a user
type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password,omitempty" bson:"password"`
}

// Register handles the full registration process for the handler
func (user *User) Register() (err error) {
	if err := user.validate(); err != nil {
		return err
	}
	return nil
}

func (user *User) validate() error {
	if user.Username == "" {
		return errors.New("no username provided")
	} else if user.Email == "" {
		return errors.New("no email provided")
	} else if user.Password == "" {
		return errors.New("no password provided")
	}
	return nil
}
