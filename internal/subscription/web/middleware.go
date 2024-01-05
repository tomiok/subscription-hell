package web

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Auth(s *session.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sess, err := s.Get(c)
		if err != nil {
			return err
		}

		userID := sess.Get("userID")
		if userID == nil {
			return errors.New("user is not authenticated")
		}

		return c.Next()
	}
}
