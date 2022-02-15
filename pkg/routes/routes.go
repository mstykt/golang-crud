package routes

import (
	"github.com/labstack/echo/v4"
	"golang-crud/pkg/controller"
)

type Routes struct {
	userController *controller.UserController
}

func NewRoutes(controller *controller.UserController) *Routes {
	return &Routes{userController: controller}
}

func (routes *Routes) InitRoutes(echo *echo.Echo) {
	echo.POST("/users", routes.userController.CreateUser)
	echo.GET("/users/:id", routes.userController.GetUser)
	echo.GET("/users", routes.userController.GetAllUser)
	echo.PUT("/users/:id", routes.userController.UpdateUser)
	echo.DELETE("/users/:id", routes.userController.DeleteUser)
}
