package http

import (
	"github.com/VQIVS/FlowForge/pkg/context"
	"github.com/VQIVS/FlowForge/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func setUserContext(c *fiber.Ctx) error {
	c.SetUserContext(context.NewAppContext(c.UserContext(), context.WithLogger(logger.NewLogger())))
	return c.Next()
}
