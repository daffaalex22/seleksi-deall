package routes

import (
	userController "github.com/daffaalex22/seleksi-deall/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteControllerList struct {
	UsersController userController.UsersController
	JWTConfig       middleware.JWTConfig
}

func (controller RouteControllerList) RouteRegister(e *echo.Echo) {

	e.Pre(middleware.RemoveTrailingSlash())

	jwtMiddleware := middleware.JWTWithConfig(controller.JWTConfig)

	//user
	e.POST("/users/login",
		controller.UsersController.UsersLogin)
	e.POST("/users",
		controller.UsersController.UsersAdd)
	e.GET("/users",
		controller.UsersController.UsersGetAll,
		jwtMiddleware)
	e.PUT("/users",
		controller.UsersController.UsersUpdate,
		jwtMiddleware)
	e.DELETE("/users/:id",
		controller.UsersController.UsersDelete,
		jwtMiddleware)
	e.GET("/users/:id",
		controller.UsersController.UsersGetById,
		jwtMiddleware)
	e.GET("/users/me",
		controller.UsersController.UsersGetMyData,
		jwtMiddleware)
}
