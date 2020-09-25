package router

import (
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/go-impatient/gaia/app/conf"
	"github.com/go-impatient/gaia/app/handler"
	sdHandle "github.com/go-impatient/gaia/app/handler/sd"
	"github.com/go-impatient/gaia/app/middleware/timer"
	"github.com/go-impatient/gaia/internal/service"
	"github.com/go-impatient/gaia/pkg/array"
)

func rootHandler(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("Welcome to api app.")
}

// NotFound creates a gin middleware for handling page not found.
func NotFound() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

// RegisterRoutes ...
func RegisterRoutes(router *fiber.App, services *service.Services) {
	// 使用中间件.
	// -------------------
	router.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, HEAD, PUT, DELETE, PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
		MaxAge:           int(12 * time.Hour),
	}))

	// Logger 中间件.
	// -------------------
	mode := conf.AppConfig.Mode
	if mode != "prod" {
		router.Use(logger.New(logger.Config{
			Next:       nil,
			Format:     "[${time}] ${status} - ${latency} - ${method} ${path}\n",
			TimeFormat: "2006-01-02 15:04:05",
			TimeZone:   "Local",
			Output:     os.Stderr,
		}))
	}

	router.Use(
		requestid.New(),
		recover.New(),
		pprof.New(),
	)

	// Timer
	//-------------------
	router.Use(timer.New(timer.Config{
		DisplayMilliseconds: true,
		DisplayHuman:        true,
	}))
	if mode == "prod" {
		router.Use(limiter.New(limiter.Config{
			Next: func(c *fiber.Ctx) bool {
				return array.StringInSlice(c.IP(), []string{"localhost"})
			},
			Max:          20,
			Duration:     30,
			Key: func(ctx *fiber.Ctx) string {
				return ctx.IP()
			},
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
					"code":    fiber.StatusTooManyRequests,
					"message": "请求太频繁了",
				})
			},
		}))
	}

	router.Get("/", rootHandler)

	// The health check handlers
	// -----------------------------
	sd := router.Group("/sd")
	{
		sd.Get("/health", sdHandle.HealthCheck)
		sd.Get("/disk", sdHandle.DiskCheck)
		sd.Get("/cpu", sdHandle.CPUCheck)
		sd.Get("/ram", sdHandle.RAMCheck)
	}

	handler.MakeUserHandler(router, services.User)

	// Custom 404 (after all routes)
	// -----------------------------
	router.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": "Resource Not Found",
		})
	})
}
