package main

import (
	"usersystem/service"
)

func main() {
	Service := service.Service{}
	Service.Init()
	Service.Run()

}
