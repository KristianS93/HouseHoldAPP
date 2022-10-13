package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
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

	// statically served files
	app.Static("/", "./public/static")

	// Middleware
	app.Use(favicon.New(favicon.Config{
		File: "./public/static/images/favicon.ico",
	}))

	// GET endpoints
	app.Get("/", Index)
	app.Get("/grocerylist", GroceryList)

	// POST endpoints
	app.Post("/additem", Additem)
	app.Post("login", Login)

	// PATCH endpoints
	app.Patch("/changeitem", ChangeItem)

	// should be changed to fetch instead of hyperlink frontend, should also change to delete method
	app.Get("/clearlist", ClearList)

	return app
}
