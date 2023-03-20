package main

import (
	"api-connect-mongodb-atlas/app/controllers"
	"api-connect-mongodb-atlas/app/routes"
	"api-connect-mongodb-atlas/pkg/configs"
	"api-connect-mongodb-atlas/pkg/middleware"
	"api-connect-mongodb-atlas/pkg/utils"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {

	configs.Godotenv()

	config := configs.FiberConfig()
	app := fiber.New(config)

	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	mongoAtlas := configs.InintMongodbAtlas()

	var (
		analyticsController = controllers.NewAnalyticsController(mongoAtlas)
		userController      = controllers.NewUserControllers(mongoAtlas)
		providerController  = controllers.NewProviderControllers(mongoAtlas)

		analyticsRoute = routes.NewAnalyticsRoute(analyticsController)
		userRoute      = routes.NewUserRoute(userController)
		providerRoute  = routes.NewProviderRoute(providerController)
	)

	analyticsRoute.AnalyticsList(app)
	userRoute.UserPropsRoute(app)
	providerRoute.ProviderPropsRoute(app)

	app.Get("/version", controllers.TestConnect)

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
