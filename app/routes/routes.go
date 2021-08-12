package routes

import (
	"shoe-store/controllers/category"
	"shoe-store/controllers/news"
	"shoe-store/controllers/users"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware      middleware.JWTConfig
	UserController     users.UserController
	NewsController     news.NewsController
	CategoryController category.CategoryController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("/register", cl.UserController.Store)
	users.GET("/token", cl.UserController.CreateToken)

	category := e.Group("category")
	category.GET("/list", cl.CategoryController.GetAll, middleware.JWTWithConfig(cl.JWTMiddleware))

	news := e.Group("news")
	news.POST("/store", cl.NewsController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))
	news.PUT("/update", cl.NewsController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))
}
