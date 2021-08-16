package main

import (
	_productController "shoe-store/controllers/product"
	_productUsecase "shoe-store/domains/product"
	_productRepo "shoe-store/drivers/databases/product"

	_brandController "shoe-store/controllers/brand"
	_brandUsecase "shoe-store/domains/brand"
	_brandRepo "shoe-store/drivers/databases/brand"

	_userController "shoe-store/controllers/users"
	_userUsecase "shoe-store/domains/users"
	_userRepo "shoe-store/drivers/databases/users"

	_transactionController "shoe-store/controllers/transaction"
	_transactionUsecase "shoe-store/domains/transaction"
	_transactionRepo "shoe-store/drivers/databases/transaction"

	_transactionItemRepo "shoe-store/drivers/databases/transactionItem"

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
		&_productRepo.Product{},
		&_brandRepo.Brand{},
		&_userRepo.Users{},
		&_transactionRepo.Transaction{},
		&_transactionItemRepo.TransactionItem{},
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

	brandRepo := _brandRepo.NewBrandRepository(db)
	brandUsecase := _brandUsecase.NewBrandUsecase(timeoutContext, brandRepo)
	brandCtrl := _brandController.NewBrandController(brandUsecase)

	productRepo := _productRepo.NewMySQLProductRepository(db)
	productUsecase := _productUsecase.NewProductUsecase(productRepo, brandUsecase, timeoutContext, iplocatorRepo)
	productCtrl := _productController.NewProductController(productUsecase)

	userRepo := _userRepo.NewMySQLUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, &configJWT, timeoutContext)
	userCtrl := _userController.NewUserController(userUsecase)

	transactionRepo := _transactionRepo.NewTransactionRepository(db)
	transactionUsecase := _transactionUsecase.NewTransactionUsecase(timeoutContext, transactionRepo)
	transactionCtrl := _transactionController.NewTransactionController(transactionUsecase)

	routesInit := _routes.ControllerList{
		JWTMiddleware:         configJWT.Init(),
		UserController:        *userCtrl,
		ProductController:     *productCtrl,
		BrandController:       *brandCtrl,
		TransactionController: *transactionCtrl,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
