package routes

import (
	"prakerja6/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) *echo.Echo{
	route := e.Group("")
	route.GET("/users", controllers.GetAllUSer)
	route.GET("/users/:id", controllers.DetailUserController)
	route.DELETE("/users/:id", controllers.DeleteUser)
	route.POST("/users", controllers.CreateUserController)
	route.POST("/books", controllers.CreateBookController)
	route.GET("/books", controllers.GetAllBook)
	route.GET("/books/:id", controllers.DetailBookController)
	route.DELETE("/books/:id", controllers.DeleteBook)
	return e
}