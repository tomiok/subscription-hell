package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/tomiok/subscription-hell/internal/subscription"
)

type SubHandler struct {
	store   *session.Store
	service *subscription.Service
}

func NewSubHandler(service *subscription.Service) *SubHandler {
	return &SubHandler{service: service}
}

func (s *SubHandler) AddSub(c *fiber.Ctx) error {
	var req subscription.SubRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	sess, err := s.store.Get(c)
	if err != nil {
		return err
	}

	userID := sess.Get("userID")
	req.UserID = userID.(uint)

	sub, err := s.service.AddSub(req)

	if err != nil {
		return err
	}

	return c.JSON(sub)
}
