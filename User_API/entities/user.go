package entities

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id"`
	FirstName       string             `json:"firstname" bson:"firstname" binding:"required"`
	LastName        string             `json:"lastname" bson:"lastname" binding:"required"`
	Age             string             `json:"age" bson:"age" binding:"required"`
	Email           string             `json:"email" bson:"email" binding:"required"`
	Password        string             `json:"password" bson:"password" binding:"required,min=8"`
	ConfirmPassword string             `json:"confirmPassword" bson:"confirmPassword,omitempty" binding:"required"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at"`
}

type SignupResponse struct {
	Name      string    `json:"firstname" bson:"firstname"`
	Email     string    `json:"email" bson:"email"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
type Login struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required,min=8"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh"`
}

// type Logout struct {
// 	Email string `json:"email" bson:"email"`
// }
// type LogoutResponse struct {
// 	Email     string    `json:"email" bson:"email"`
// 	LoggedOut time.Time `json:"logged_out" bson:"logged_out"`
// }

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	User_type  string
	jwt.StandardClaims
}
