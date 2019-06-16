package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is the type of a user
type User struct {
	ID       primitive.ObjectID `jsin:"id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"-" bson:"password"`
}
