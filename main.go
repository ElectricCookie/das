package main

//go:generate go-bindata -pkg assets -o assets/assets.go assets/files

import "github.com/ElectricCookie/das-cms/user"
import "github.com/ElectricCookie/das-cms/routes"
import "github.com/ElectricCookie/das-cms/db"
import "github.com/ElectricCookie/das-cms/i18n"
import "github.com/derekparker/delve/pkg/config"
import "github.com/ElectricCookie/das-cms/configLoader"

func main() {
	config.LoadConfig()
	i18n.LoadTranslations()
	db.Connect()
	routes.CreateRouter()
	user.RegisterNamespace()

	routes.Router.Run(configLoader.GetConfig().Interface)

}
