package main

import (
	"github.com/go-impatient/gaia/pkg/http/fiberhttp"
	"github.com/gofiber/fiber/v2"
)

func main()  {
	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	//app.Get("/", func(c *fiber.Ctx) error {
	//	return c.SendString("I'm in prefork mode 🚀")
	//})

	server := fiberhttp.NewServer(fiberhttp.Addr(":4000"), fiberhttp.App(app))
	server.Router().Get("/", func(c *fiber.Ctx) error {
		return c.SendString("I'm in prefork mode 🚀")
	})
	server.Serve()



	// log.Println(server.Run(":3000"))
}
