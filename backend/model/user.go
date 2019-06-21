package model

import (
	"errors"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var (
	emailRegex = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
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
	if valid, err := user.validate(); !valid {
		return err
	}
	if err = user.hashPassword(); err != nil {
		return err
	}
	return nil
}

func (user *User) hashPassword() error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return errors.New("can't hash password")
	}
	user.Password = string(password)
	return nil
}

func (user *User) validate() (bool, error) {
	if valid, err := user.validateUsername(); !valid {
		return valid, err
	} else if valid, err := user.validateEmail(); !valid {
		return valid, err
	} else if valid, err := user.validatePassword(); !valid {
		return valid, err
	}
	return true, nil
}

func (user *User) validateUsername() (bool, error) {
	if user.Username == "" {
		return false, errors.New("no username provided")
	} else if match, _ := regexp.MatchString("^[a-zA-Z-_1-9]+$", user.Username); !match {
		return false, errors.New("username has wrong pattern")
	}
	return true, nil
}

func (user *User) validateEmail() (bool, error) {
	if user.Email == "" {
		return false, errors.New("no email address provided")
	} else if match, _ := regexp.MatchString(emailRegex, user.Email); !match {
		return false, errors.New("email has wrong pattern")
	}
	return true, nil
}

func (user *User) validatePassword() (bool, error) {
	if user.Password == "" {
		return false, errors.New("no password provided")
	} else if match, _ := regexp.MatchString("^\\S+$", user.Password); !match {
		return false, errors.New("password has whitespace")
	}
	return true, nil
}
