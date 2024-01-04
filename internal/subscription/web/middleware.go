package web

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func Auth(c *fiber.Ctx) error {
	log.Print("this is called")
	return c.Next()
}
