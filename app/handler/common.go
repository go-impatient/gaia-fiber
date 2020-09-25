package handler

import (
	"github.com/gofiber/fiber/v2"
)

// handler is a wrapper that allows the the server route functions to return
// an error. This is useful, because otherwise you would have to do the call
// to the Next handler call on each error. Ain't no body got time for that
func handlerWrapper(h fiber.Handler) fiber.Handler{
	return func(c *fiber.Ctx) error {
		if err := h(c); err != nil {
			c.Next()
			return err
		}
		return nil
	}
}