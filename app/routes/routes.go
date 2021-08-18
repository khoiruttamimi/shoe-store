package routes

import (
	"shoe-store/controllers/brand"
	"shoe-store/controllers/product"
	"shoe-store/controllers/transaction"
	"shoe-store/controllers/users"
	"shoe-store/controllers/wishlist"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware         middleware.JWTConfig
	UserController        users.UserController
	ProductController     product.ProductController
	BrandController       brand.BrandController
	TransactionController transaction.TransactionController
	WishlistController    wishlist.WishlistController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	jwtAuth := middleware.JWTWithConfig(cl.JWTMiddleware)
	api := e.Group("v1/api/")
	users := api.Group("users")
	users.POST("/register", cl.UserController.RegisterUser)
	users.POST("/register-admin", cl.UserController.RegisterAdmin)
	users.POST("/login", cl.UserController.Login)
	users.GET("/profile", cl.UserController.GetProfile, jwtAuth)
	users.PUT("/profile", cl.UserController.UpdateProfile, jwtAuth)

	brand := api.Group("brands", jwtAuth)
	brand.GET("", cl.BrandController.GetAll)
	//hanya admin
	brand.POST("", cl.BrandController.Store)

	product := api.Group("products", jwtAuth)
	product.GET("", cl.ProductController.GetAll)
	product.GET("/:id", cl.ProductController.GetByID)
	//hanya admin
	product.POST("", cl.ProductController.Store)
	product.PUT("/:id", cl.ProductController.Update)

	transaction := api.Group("transactions", jwtAuth)
	transaction.POST("", cl.TransactionController.Shopping)
	transaction.GET("", cl.TransactionController.GetAll)
	transaction.GET("/:id", cl.TransactionController.GetByID)

	wishlist := api.Group("wishlist", jwtAuth)
	wishlist.POST("", cl.WishlistController.AddWishlist)
	wishlist.GET("", cl.WishlistController.GetAll)
}
