package main

import (
	"log"
	"server/web"
)

func main() {
	log.Fatal(web.GetApp().Listen("localhost:8888"))
}
