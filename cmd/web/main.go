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

	viewsRouter := app.Group("v")
	viewsRouter.Get("/sign-up", userWeb.SignUpView)
	viewsRouter.Get("/login", userWeb.LoginView)

	bizRouter := app.Group("b")
	bizRouter.Post("/sign-up", userWeb.Signup)
	bizRouter.Get("/", userWeb.HomeView)

	log.Fatal(app.Listen(":3000"))
}
