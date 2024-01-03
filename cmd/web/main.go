package main

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/tomiok/subscription-hell/internal/subscription"
	"github.com/tomiok/subscription-hell/internal/subscription/web"
	"github.com/tomiok/subscription-hell/pkg/database"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type deps struct {
	store        *session.Store
	userService  *subscription.UserService
	usersHandler *web.User
}

func main() {
	sess := session.New()
	userWeb := web.NewWebUser(nil, sess)

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

func init() {
	database.InitDB("localhost", "subs", "subs", "subs", 5432)
}

func newDeps() *deps {

	return &deps{
		store:       nil,
		userService: nil,
	}
}
