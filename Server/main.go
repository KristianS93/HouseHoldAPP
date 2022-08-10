package main

import "server/web"

func main() {
	Server := web.Server{}
	Server.Init()
	Server.Run()
}
