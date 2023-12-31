package main

import (
	"github.com/tomiok/subscription-hell/internal/subscription/web"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	userWeb := web.NewWebUser()

	// Create a new engine
	engine := html.New("./views", ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})
	app.Static("/static/css", "./static/css")

	app.Get("/sign-up", userWeb.SignUp)
	app.Get("/login", userWeb.SignUp)

	log.Fatal(app.Listen(":3000"))
}
