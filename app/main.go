package main

import (
	_newsUsecase "shoe-store/businesses/news"
	_newsController "shoe-store/controllers/news"
	_newsRepo "shoe-store/drivers/databases/news"

	_categoryUsecase "shoe-store/businesses/category"
	_categoryController "shoe-store/controllers/category"
	_categoryRepo "shoe-store/drivers/databases/category"

	_userUsecase "shoe-store/businesses/users"
	_userController "shoe-store/controllers/users"
	_userRepo "shoe-store/drivers/databases/users"

	_dbDriver "shoe-store/drivers/mysql"

	_ipLocatorDriver "shoe-store/drivers/thirdparties/iplocator"

	_middleware "shoe-store/app/middleware"
	_routes "shoe-store/app/routes"

	"log"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_newsRepo.News{},
		&_categoryRepo.Category{},
		&_userRepo.Users{},
	)
}

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configDB.InitialDB()
	dbMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()

	iplocatorRepo := _ipLocatorDriver.NewIPLocator()

	categoryRepo := _categoryRepo.NewCategoryRepository(db)
	categoryUsecase := _categoryUsecase.NewCategoryUsecase(timeoutContext, categoryRepo)
	categoryCtrl := _categoryController.NewCategoryController(categoryUsecase)

	newsRepo := _newsRepo.NewMySQLNewsRepository(db)
	newsUsecase := _newsUsecase.NewNewsUsecase(newsRepo, categoryUsecase, timeoutContext, iplocatorRepo)
	newsCtrl := _newsController.NewNewsController(newsUsecase)

	userRepo := _userRepo.NewMySQLUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, &configJWT, timeoutContext)
	userCtrl := _userController.NewUserController(userUsecase)

	routesInit := _routes.ControllerList{
		JWTMiddleware:      configJWT.Init(),
		UserController:     *userCtrl,
		NewsController:     *newsCtrl,
		CategoryController: *categoryCtrl,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
