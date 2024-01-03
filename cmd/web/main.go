package main

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/tomiok/subscription-hell/internal/subscription"
	"github.com/tomiok/subscription-hell/internal/subscription/storage"
	"github.com/tomiok/subscription-hell/internal/subscription/web"
	"github.com/tomiok/subscription-hell/pkg/database"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type deps struct {
	usersHandler *web.User
}

func main() {
	d := newDeps()

	// Create a new engine
	engine := html.New("./views", ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})
	app.Static("/static/css", "./static/css")

	viewsRouter := app.Group("v")
	viewsRouter.Get("/sign-up", d.usersHandler.SignUpView)
	viewsRouter.Get("/login", d.usersHandler.LoginView)

	bizRouter := app.Group("b")
	bizRouter.Post("/sign-up", d.usersHandler.Signup)
	bizRouter.Get("/", d.usersHandler.HomeView)

	log.Fatal(app.Listen(":3000"))
}

func init() {
	database.InitDB("localhost", "subs", "subs", "subs", 5432)

	if err := database.Migrate(subscription.User{}, subscription.Sub{}); err != nil {
		log.Fatal(err)
	}
}

func newDeps() *deps {
	store := session.New()
	userStorage := storage.NewUserStorage()
	userService := subscription.NewUserService("secret-change-me", userStorage)
	usersHandler := web.NewWebUser(userService, store)

	return &deps{
		usersHandler: usersHandler,
	}
}
