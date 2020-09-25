package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/go-impatient/gaia/internal/service"
)

// UserHandler ...
type userHandler struct {
	userService service.UserService
}

// Login 用户登录
func (handler *userHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		handler.userService.Login(c.Context(), "", "")
		c.Status(http.StatusOK).JSON(&fiber.Map{
			"success": true,
			"message": "登录成功",
		})
		return nil
	}
}

// Register 注册
func (handler *userHandler) Register() fiber.Handler{
	return func(c *fiber.Ctx) error {
		handler.userService.Register(c.Context(), "", "", "")
		c.JSON(&fiber.Map{
			"success": true,
			"message": "注册成功",
		})
		return nil
	}
}

func MakeUserHandler(r *fiber.App, srv service.UserService) {
	handler := &userHandler{userService: srv}

	userGroup := r.Group("/user")
	{
		userGroup.Post("/login", handler.Login())
		userGroup.Post("/register", handler.Register())
	}
}
