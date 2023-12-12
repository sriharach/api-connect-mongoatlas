package routes

import (
	"api-connect-mongodb-atlas/app/controllers"

	"github.com/gofiber/fiber/v2"
)

type IProviderRoute interface {
	ProviderPropsRoute(a *fiber.App)
}

type ProviderRouteTool struct {
	ProviderInterface controllers.IProviders
}

func NewProviderRoute(ac controllers.IProviders) IProviderRoute {
	return &ProviderRouteTool{
		ProviderInterface: ac,
	}
}

func (pr *ProviderRouteTool) ProviderPropsRoute(a *fiber.App) {
	group := a.Group("/api")
	group.Post("/register", pr.ProviderInterface.RegisterAccount)
	group.Post("/login", pr.ProviderInterface.Login)
	group.Get("/auth/url", pr.ProviderInterface.Oauth2)
	group.Get("/oauth/callback", pr.ProviderInterface.OauthCallback)
	group.Get("/logout", pr.ProviderInterface.Logout)
}
