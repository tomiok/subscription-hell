package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/tomiok/subscription-hell/internal/subscription"
)

type User struct {
	*subscription.UserService
	store *session.Store
}

func NewWebUser(
	service *subscription.UserService,
	store *session.Store,
) *User {
	return &User{
		UserService: service,
		store:       store,
	}
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
	sess, err := u.store.Get(c)
	if err != nil {
		return err
	}

	userID := sess.Get("userID")
	return c.Render("index", fiber.Map{
		"userID": userID,
	})
}

/*
logic
*/

// Signup
func (u *User) Signup(c *fiber.Ctx) error {
	nick := c.FormValue("nickname")
	pass := c.FormValue("password")

	user, err := u.UserService.CreateUser(nick, pass)
	if err != nil {
		return err
	}

	sess, err := u.store.Get(c)
	if err != nil {
		return err
	}

	sess.Set("userID", user.ID)

	return c.Redirect("/b")
}
