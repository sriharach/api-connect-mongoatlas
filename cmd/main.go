package main

import (
	"api-connect-mongodb-atlas/app/controllers"
	"api-connect-mongodb-atlas/app/routes"
	"api-connect-mongodb-atlas/pkg/configs"
	"api-connect-mongodb-atlas/pkg/middleware"
	"api-connect-mongodb-atlas/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {

	configs.Godotenv()

	config := configs.FiberConfig()
	app := fiber.New(config)

	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	mongoAtlas := configs.InintMongodbAtlas()

	var (
		userController     = controllers.NewUserControllers(mongoAtlas)
		providerController = controllers.NewProviderControllers(mongoAtlas)

		userRoute     = routes.NewUserRoute(userController)
		providerRoute = routes.NewProviderRoute(providerController)
	)

	userRoute.UserPropsRoute(app)
	providerRoute.ProviderPropsRoute(app)

	app.Get("/version", controllers.TestConnect)

	utils.StartServer(app)
}
