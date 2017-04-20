package configLoader

import (
	"encoding/json"
	"io/ioutil"

	validator "gopkg.in/go-playground/validator.v9"
)

// Configuration describes the main CMS configuration structure
type Configuration struct {
	URL          string `json:"url" validate:"required,url"`
	SMTPURL      string `json:"smtpUrl" validate:"required,url"`
	SMTPPort     int    `json:"smtpPort" validate:"required"`
	SMTPUser     string `json:"smtpUser" validate:"required"`
	SMTPPassword string `json:"smtpPassword" validate:"required"`
	SMTPFrom     string `json:"stmpFrom" validate:"required,email"`
	JWTSecret    string `json:"jwtSecret" validate:"required"`
}

var config Configuration

// LoadConfiguration loads the config.json file
func LoadConfiguration() {

	validate := validator.New()

	data, err := ioutil.ReadFile("config.json")

	if err != nil {
		panic("No config file found!")
	}

	config = Configuration{}

	if err := json.Unmarshal(data, &config); err != nil {
		panic("Invalid configuration file!")
	}

	if err := validate.Struct(&config); err != nil {
		panic(err)
	}

}

// GetConfig returns the current configuration
func GetConfig() *Configuration {
	return &config
}
