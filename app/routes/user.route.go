package routes

import (
	"api-connect-mongodb-atlas/app/controllers"
	"api-connect-mongodb-atlas/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

type IUserRoute interface {
	UserPropsRoute(a *fiber.App)
}

type UserRouteTool struct {
	UserInterface controllers.IuserController
}

func NewUserRoute(ac controllers.IuserController) IUserRoute {
	return &UserRouteTool{
		UserInterface: ac,
	}
}

func (ct *UserRouteTool) UserPropsRoute(a *fiber.App) {
	group := a.Group("/api/v1", middleware.DeserializeUser)
	group.Get("/user/profile", ct.UserInterface.GetUserAccount)
	group.Get("/users", ct.UserInterface.GetUsersAccount)
	group.Put("/user/update/:_id", ct.UserInterface.EditUserAccount)
	group.Delete("/user/remove/:_id", ct.UserInterface.DeleteUserAccount)
}
