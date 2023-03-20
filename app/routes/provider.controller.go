package routes

import (
	"api-connect-mongodb-atlas/app/controllers"

	"github.com/gofiber/fiber/v2"
)

type IProviderRoute interface {
	ProviderPropsRoute(a *fiber.App)
}

type ProviderRouteTool struct {
	UserInterface controllers.IProviders
}

func NewProviderRoute(ac controllers.IProviders) IProviderRoute {
	return &ProviderRouteTool{
		UserInterface: ac,
	}
}

func (pr *ProviderRouteTool) ProviderPropsRoute(a *fiber.App) {
	group := a.Group("/api")
	group.Post("/login", pr.UserInterface.Login)
}
