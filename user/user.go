package user

import (
	"github.com/ElectricCookie/das-cms/routes"

	validator "gopkg.in/go-playground/validator.v9"
)

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
