package timer

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Config defines the configuration for middleware.
type Config struct {
	// DisplaySeconds indicates the process time in seconds.
	//
	// Optional. Default value false.
	DisplaySeconds bool

	// DisplayMilliseconds indicates the process time in milliseconds.
	//
	// Optional. Default value false.
	DisplayMilliseconds bool

	// DisplayHuman indicates the process time in human format.
	//
	// Optional. Default value true.
	DisplayHuman bool

	// Prefix indicates prefix for header name.
	//
	// Optional. Default value "X-Process-Time".
	Prefix string
}

// ConfigDefault is the default configuration.
var ConfigDefault = Config{
	DisplaySeconds:      false,
	DisplayMilliseconds: false,
	DisplayHuman:        true,
	Prefix:              "X-Process-Time",
}

// New creates a new middleware handler
func New(config ...Config) func(*fiber.Ctx) error {
	// Default configuration
	cfg := ConfigDefault

	// Override configuration if provided
	if len(config) > 0 {
		cfg = config[0]

		if cfg.Prefix == "" {
			cfg.Prefix = ConfigDefault.Prefix
		}
	}

	return func(c *fiber.Ctx) error {
		now := time.Now()

		c.Next()

		duration := time.Since(now)
		if cfg.DisplayHuman {
			c.Set(cfg.Prefix, fmt.Sprintf("%v", duration))
		}

		if cfg.DisplayMilliseconds {
			c.Set(cfg.Prefix+"-ms", fmt.Sprintf("%v", duration.Milliseconds()))
		}

		if cfg.DisplaySeconds {
			c.Set(cfg.Prefix+"-sec", fmt.Sprintf("%v", duration.Seconds()))
		}

		return nil
	}
}
