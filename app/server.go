package app

import (
	"github.com/roby-aw/go-clean-architecture-hexagonal/api"
	"github.com/roby-aw/go-clean-architecture-hexagonal/app/modules"
	"github.com/roby-aw/go-clean-architecture-hexagonal/config"
	"github.com/roby-aw/go-clean-architecture-hexagonal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	_ "github.com/roby-aw/go-clean-architecture-hexagonal/docs"

	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func Run(config *config.AppConfig, dbCon *utils.DatabaseConnection) (*fiber.App, string) {
	app := fiber.New()

	app.Use(recover.New())

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] [${ips}:${port}] ${status} - ${latency} ${method} ${url}\n",
		TimeFormat: "2 Jan 2006 15:04:05",
	}))
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "API IS UP!!!!",
		})
	})

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	controller := modules.RegistrationModules(dbCon, config)
	api.RegistrationPath(app, controller)
	return app, config.App.Port
}
