package main

import "grocerylist/service"

func main() {
	//initialize server from server/web package
	Server := service.Server{}

	//Initialize server settings
	Server.Init()
	//Run server/service instance.
	Server.Run()
}
