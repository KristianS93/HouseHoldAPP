package main

import (
	"mealplanner/database"
	"mealplanner/service"
)

func main() {
	//initialize server from server/web package
	Server := service.Server{}

	database.Connect()

	// meal.DBConnection()
	// meal.DB.AutoMigrate(meal.User{})
	//Initialize server settings
	Server.Init()
	//Run server/service instance.
	Server.Run(":5005")
}
