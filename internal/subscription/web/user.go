package web

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
}

func NewWebUser() *User {
	return &User{}
}

/*
views
*/

// SignUpView shows up sign view.
func (u *User) SignUpView(c *fiber.Ctx) error {
	return c.Render("signup", fiber.Map{})
}

// LoginView shows up login view.
func (u *User) LoginView(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{})
}

// HomeView shows up home view (index).
func (u *User) HomeView(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

/*
logic
*/

// Signup
func (u *User) Signup(c *fiber.Ctx) error {
	return c.Redirect("/b/")
}
