package main

import (
	"api-connect-mongodb-atlas/app/controllers"
	"api-connect-mongodb-atlas/app/routes"
	"api-connect-mongodb-atlas/pkg/configs"
	"api-connect-mongodb-atlas/pkg/middleware"
	"api-connect-mongodb-atlas/pkg/utils"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}

func main() {

	configs.Godotenv()

	config := configs.FiberConfig()
	app := fiber.New(config)

	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	mongoAtlas := configs.InintMongodbAtlas()

	var (
		analyticsController = controllers.NewAnalyticsController(mongoAtlas)
		userController      = controllers.NewUserControllers(mongoAtlas)

		analyticsRoute = routes.NewAnalyticsRoute(analyticsController)
		userRoute      = routes.NewUserRoute(userController)
	)

	analyticsRoute.AnalyticsList(app)
	userRoute.UserPropsRoute(app)

	app.Get("/version", controllers.TestConnect)

	// Start server (with or without graceful shutdown).
	app.Server()
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
