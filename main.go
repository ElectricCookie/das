package main

//go:generate go-bindata -pkg assets -o assets/assets.go assets/files

import "github.com/ElectricCookie/das-cms/user"
import "github.com/ElectricCookie/das-cms/routes"
import "github.com/ElectricCookie/das-cms/db"
import "github.com/ElectricCookie/das-cms/i18n"
import "github.com/derekparker/delve/pkg/config"
import "github.com/ElectricCookie/das-cms/configLoader"
import "github.com/ElectricCookie/das-cms/mongo"

func main() {
	config.LoadConfig()

	i18n.LoadTranslations()
	mongoAdapter := mongo.Adapter{}
	mongoAdapter.Connect(configLoader.GetConfig().DBHost)
	db.SetConnector(mongoAdapter)
	routes.CreateRouter()
	user.RegisterNamespace()

	routes.Router.Run(configLoader.GetConfig().Interface)

}
