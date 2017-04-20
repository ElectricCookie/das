package user

import (
	"github.com/ElectricCookie/das-cms/db"
	"github.com/ElectricCookie/das-cms/routes"

	validator "gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2/bson"
)

// User describes a user in the system
type User struct {
	ID               string `json:"id"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"LastName"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Email            string `json:"email"`
	Language         string `json:"language"`
	EmailVerified    bool   `json:"emailVerified"`
	EmailVerifyToken string `json:"emailVerifyToken"`
	Salt             string `json:"salt"`
	Registered       int64  `json:"registered"`
	LastLogin        int64  `json:"lastLogin"`
}

// ForgotPassword is the required information required to request a reset password email
type ForgotPassword struct {
	Email string `json:"email"`
}

// ResetPasswordData is the required information to reset the password
type ResetPasswordData struct {
	Token       string `json:"token"`
	NewPassword string `json:"newPassword"`
}

var validate *validator.Validate

// RegisterNamespace registers all functions for users
func RegisterNamespace() {

	validate = validator.New()

	routes.Router.POST("/login", loginHandler)

	routes.Router.POST("/register", registerHandler)

}

// GetUserByID retrieves a user by ID
func GetUserByID(userID string) (*User, error) {
	var user User
	err := db.GetDb().C("users").Find(bson.M{
		"_id": userID,
	}).One(user)

	return &user, err

}
