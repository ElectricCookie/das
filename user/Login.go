package user

import (
	"net/http"

	"github.com/ElectricCookie/das-cms/db"
	"github.com/ElectricCookie/das-cms/routes"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/scrypt"
)

// LoginData is the required information for logging in
type LoginData struct {
	UsernameOrEmail string `json:"id" binding:"required"`
	Password        string `json:"password" binding:"required"`
	RememberMe      bool   `json:"rememberMe" binding:"required"`
}

func loginHandler(c *gin.Context) {

	params := LoginData{}

	if c.BindJSON(&params) != nil {
		c.JSON(http.StatusBadRequest, routes.InvalidParams)
		return
	}

	token, err := Login(params.UsernameOrEmail, params.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	expires := 2147483647

	if !params.RememberMe {
		expires = 3600
	}

	c.SetCookie("RefreshToken", *token, expires, "/logout", "", true, true)
	c.SetCookie("RefreshToken", *token, expires, "/acess-token", "", true, true)

}

// Login a user
func Login(usernameOrEmail string, password string) (*string, *routes.APIError) {

	user, err := db.GetDb().GetUserByUsernameOrEmail(usernameOrEmail)

	if err != nil {
		return nil, &routes.APIError{
			ErrorCode:   "unknownUser",
			Description: "The user requested was not found",
		}
	}

	dk, cryptError := scrypt.Key([]byte(password), []byte(user.Salt), 16384, 8, 1, 32)

	if cryptError != nil {
		return nil, &routes.InternalError
	}

	if string(dk) != user.Password {

		return nil, &routes.APIError{
			ErrorCode:   "wrongPassword",
			Description: "The passsword provided was not correct",
		}

	}

	token, tokenError := GenerateRefreshToken(user.ID)

	if tokenError != nil {
		return nil, &routes.InternalError
	}

	return token, nil

}
