package web

import "github.com/gofiber/fiber/v2"

type User struct {
}

func NewWebUser() *User {
	return &User{}
}

func (u *User) SignUpView(c *fiber.Ctx) error {
	return c.Render("signup", fiber.Map{})
}

func (u *User) LoginView(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{})
}
