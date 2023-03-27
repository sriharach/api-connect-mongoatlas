package main

import (
	"api-connect-mongodb-atlas/app/controllers"
	"api-connect-mongodb-atlas/app/routes"
	"api-connect-mongodb-atlas/pkg/configs"
	"api-connect-mongodb-atlas/pkg/middleware"
	"api-connect-mongodb-atlas/pkg/utils"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Easier to get running with CORS. Thanks for help @Vindexus and @erkie
var allowOriginFunc = func(r *http.Request) bool {
	return true
}

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

	app.Post("/cookie", func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name:     "cookie",
			Value:    "cookie",
			Expires:  time.Now().Add(24 * time.Hour),
			HTTPOnly: true,
			SameSite: "lax",
		})
		return nil
	})

	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
