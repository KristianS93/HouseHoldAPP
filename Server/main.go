package main

import (
	"log"
	"server/web"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./public/templates", ".html")

	app := fiber.New(fiber.Config{
		StrictRouting:     true,
		CaseSensitive:     true,
		AppName:           "HouseHoldApp x Fiber v0.1",
		Views:             engine,
		EnablePrintRoutes: true,
	})
	app.Static("/", "./public/static")

	app.Get("/", web.Index)
	app.Get("/grocerylist", web.GroceryList)

	app.Post("/additem", web.Additem)

	app.Patch("changeitem", web.ChangeItem)

	log.Fatal(app.Listen("localhost:8888"))
}
