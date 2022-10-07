package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func GetApp() *fiber.App {
	engine := html.New("./public/templates", ".html")

	app := fiber.New(fiber.Config{
		AppName:           "HouseHoldApp x Fiber v0.1",
		StrictRouting:     true,
		CaseSensitive:     true,
		Views:             engine,
		EnablePrintRoutes: true,
		ErrorHandler:      ErrorHandler,
	})
	app.Static("/", "./public/static")

	app.Get("/", Index)
	app.Get("/grocerylist", GroceryList)

	app.Post("/additem", Additem)
	app.Post("login", Login)

	app.Patch("/changeitem", ChangeItem)

	// should be changed to fetch instead of hyperlink frontend, should also change to delete method
	app.Get("/clearlist", ClearList)

	return app
}
