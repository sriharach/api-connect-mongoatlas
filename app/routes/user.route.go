package routes

import (
	"api-connect-mongodb-atlas/app/controllers"

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
	group := a.Group("/api/v1")
	group.Post("/register", ct.UserInterface.RegisterAccount)
	group.Get("/user/profile", ct.UserInterface.GetUserAccount)
	group.Get("/users", ct.UserInterface.GetUsersAccount)
}
