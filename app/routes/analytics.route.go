package routes

import (
	"api-connect-mongodb-atlas/app/controllers"

	"github.com/gofiber/fiber/v2"
)

type IAnalyticsRoute interface {
	AnalyticsList(a *fiber.App)
}

type AnalyticsRouteTool struct {
	AnalyticsInterface controllers.IAnalyticsController
}

func NewAnalyticsRoute(ac controllers.IAnalyticsController) IAnalyticsRoute {
	return &AnalyticsRouteTool{
		AnalyticsInterface: ac,
	}
}

func (ct *AnalyticsRouteTool) AnalyticsList(a *fiber.App) {
	group := a.Group("/api/v1")
	group.Get("/analytics", ct.AnalyticsInterface.GetAccount)
}
