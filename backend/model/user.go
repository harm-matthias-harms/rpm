package model

import (
	"errors"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var (
	emailRegex = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	// ExerciseAdminRole is the role of the author of an exercise
	ExerciseAdminRole = "admin"
	// TrainerRole is the role for a trainer
	TrainerRole = "trainer"
	// RolePlayManagerRole is the role for an role play manager
	RolePlayManagerRole = "role play manager"
	// MakeUpCenterRole is the role for a make up center
	MakeUpCenterRole = "make-up center"
)

// User is the type of a user
type User struct {
	ID       primitive.ObjectID `json:"id" form:"id" bson:"_id,omitempty"`
	Username string             `json:"username" form:"username" bson:"username"`
	Email    string             `json:"email" form:"email" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" form:"password" bson:"password"`
	Code     string             `json:"code,omitempty" form:"code" bson:"code,omitempty"`
	Roles    []UserRole         `json:"roles" form:"roles" bson:"roles"`
}

// LimitedUser describes a User with none vulnerable information
type LimitedUser struct {
	ID       primitive.ObjectID `json:"id" form:"id" bson:"_id" query:"id"`
	Username string             `json:"username" form:"password" bson:"username" query:"username"`
	Email    string             `json:"email" form:"email" bson:"email" query:"email"`
	Code     string             `json:"code" form:"code" bson:"code,omitempty"`
}

// UserQuery describes a query format or the user
type UserQuery struct {
	Username string `query:"username"`
	Email    string `query:"email"`
	Page     int    `query:"page"`
	PageSize int    `query:"limit"`
}

// Trainer describes a trainer role in an exercise
type Trainer struct {
	*UserRole
	Team Team `json:"team" bson:"team"`
}

// UserRole describes a management role in an exercise
type UserRole struct {
	Role     string         `json:"role" bson:"role"`
	Exercise *ExerciseShort `json:"exercise" bson:"exercise"`
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
	// email is only needed if its is not an autogenerated user
	if user.Code == "" {
		if user.Email == "" {
			return errors.New("no email address provided")
		} else if match, _ := regexp.MatchString(emailRegex, user.Email); !match {
			return errors.New("email has wrong pattern")
		}
	}
	return nil
}

func (user *User) validatePassword() error {
	if user.Code == "" {
		if user.Password == "" {
			return errors.New("no password provided")
		} else if match, _ := regexp.MatchString("^\\S+$", user.Password); !match {
			return errors.New("password has whitespace")
		}
	}
	return nil
}
