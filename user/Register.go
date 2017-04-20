package user

import (
	"net/http"

	"golang.org/x/crypto/scrypt"

	"gopkg.in/mgo.v2/bson"

	"time"

	"fmt"

	"github.com/ElectricCookie/das-cms/db"
	"github.com/ElectricCookie/das-cms/routes"
	"github.com/gin-gonic/gin"
)

// RegisterData is the required information for registering
type RegisterData struct {
	Username  string `json:"username" validate:"required,min=3,max=16"`
	Password  string `json:"password" validate:"required,min=3"`
	Email     string `json:"email" validate:"required,email"`
	Language  string `json:"language" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

func registerHandler(c *gin.Context) {
	params := RegisterData{}

	if err := c.BindJSON(&params); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, routes.InvalidParams)
		return
	}

	if err := validate.Struct(&params); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, routes.InvalidParams)
		return
	}

	if err := Register(params); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, routes.EmptyReply)

}

// Register a new user account
func Register(params RegisterData) *routes.APIError {

	user := User{}

	err := db.GetDb().C("users").Find(bson.M{
		"username": params.Username,
	}).One(&user)

	if err == nil {
		fmt.Println(err)
		return &routes.APIError{
			ErrorCode:   "usernameTaken",
			Description: "The desired username is already in use",
		}
	}

	err = db.GetDb().C("users").Find(bson.M{
		"email": params.Email,
	}).One(&user)

	if err == nil {
		return &routes.APIError{
			ErrorCode:   "emailTaken",
			Description: "The desired email is already in use",
		}
	}

	// Generate salt

	salt, saltGenerationErr := generateRandomString(64)

	if saltGenerationErr != nil {
		return &routes.InternalError
	}

	emailToken, emailTokenGenerationErr := generateRandomString(64)

	if emailTokenGenerationErr != nil {
		return &routes.InternalError
	}

	passwordHash, cryptError := scrypt.Key([]byte(params.Password), []byte(salt), 16384, 8, 1, 32)

	if cryptError != nil {
		return &routes.InternalError
	}

	// Ready to insert

	newUser := User{
		Username:         params.Username,
		Email:            params.Email,
		FirstName:        params.FirstName,
		LastName:         params.LastName,
		Password:         string(passwordHash),
		Salt:             salt,
		Language:         params.Language,
		EmailVerified:    false,
		EmailVerifyToken: emailToken,
		LastLogin:        0,
		Registered:       time.Now().Unix(),
	}

	db.GetDb().C("users").Insert(newUser)

	if err := SendRegistrationEmail(newUser); err != nil {
		fmt.Println(err)
		return &routes.InternalError
	}

	return nil
}
