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

// Authenticate checks if the user is matching to the given password.
func (user *User) Authenticate(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

// HashPassword hashes the users password for storing
func (user *User) HashPassword() error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return errors.New("can't hash password")
	}
	user.Password = string(password)
	return nil
}

// Validate validates alle field of the user have a valid form
func (user *User) Validate() (err error) {
	if err = user.validateUsername(); err != nil {
		return
	} else if err = user.validateEmail(); err != nil {
		return
	} else if err = user.validatePassword(); err != nil {
		return
	}
	return nil
}

func (user *User) validateUsername() error {
	if user.Username == "" {
		return errors.New("no username provided")
	} else if match, _ := regexp.MatchString("^[a-z1-9\\.]+$", user.Username); !match {
		return errors.New("username has wrong pattern")
	}
	return nil
}

func (user *User) validateEmail() error {
	if user.Email == "" {
		return errors.New("no email address provided")
	} else if match, _ := regexp.MatchString(emailRegex, user.Email); !match {
		return errors.New("email has wrong pattern")
	}
	return nil
}

func (user *User) validatePassword() error {
	if user.Password == "" {
		return errors.New("no password provided")
	} else if match, _ := regexp.MatchString("^\\S+$", user.Password); !match {
		return errors.New("password has whitespace")
	}
	return nil
}
