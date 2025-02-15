package http

import (
	"fmt"

	"github.com/VQIVS/FlowForge/app"
	"github.com/VQIVS/FlowForge/config"
	"github.com/gofiber/fiber/v2"
)

func Run(appContainer app.App, cfg config.ServerConfig) error {
	router := fiber.New()
	api := router.Group("/api/v1", setUserContext)

	registerAuthApi(appContainer, cfg, api)

	return router.Listen(fmt.Sprintf(":%d", cfg.HttpPort))
}

func registerAuthApi(appContainer app.App, cfg config.ServerConfig, router fiber.Router) {
	// Register auth api here
}
