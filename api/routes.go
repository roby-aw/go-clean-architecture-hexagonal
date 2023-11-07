package api

import (
	"github.com/roby-aw/go-clean-architecture-hexagonal/api/user"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	UserController *user.Controller
}

func RegistrationPath(e *fiber.App, controller Controller) {
	route := e.Group("/v1")
	routeUser := route.Group("/user")
	routeUser.Post("/login", controller.UserController.Login)
	routeUser.Post("/register", controller.UserController.Register)
	routeUser.Delete("/logout", controller.UserController.Logout)
	routeUser.Get("/me", controller.UserController.GetMe)
}
