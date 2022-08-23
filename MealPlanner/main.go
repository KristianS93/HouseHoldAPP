package main

import (
	"mealplanner/service"
	"mealplanner/service/meal"
)

func main() {
	//initialize server from server/web package
	Server := service.Server{}
	meal.DBConnection()
	// meal.DBConnection()
	// meal.DB.AutoMigrate(meal.User{})
	//Initialize server settings
	Server.Init()
	//Run server/service instance.
	Server.Run(":5005")
}
