package main

import (
	"fmt"
	"log"
	"os"
	"time"

	usersUseCase "github.com/daffaalex22/seleksi-deall/business/users"
	usersController "github.com/daffaalex22/seleksi-deall/controllers/users"
	usersRepo "github.com/daffaalex22/seleksi-deall/drivers/database/users"
	"github.com/daffaalex22/seleksi-deall/helper/encrypt"
	uuid "github.com/satori/go.uuid"

	"github.com/daffaalex22/seleksi-deall/app/middlewares"
	"github.com/daffaalex22/seleksi-deall/app/routes"
	"github.com/daffaalex22/seleksi-deall/drivers/database/mysql"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("This Services RUN on DEBUG Mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(&usersRepo.Users{})

	hashedPassword, err := encrypt.Hash("deall123")
	if err != nil {
		fmt.Println("Error hashing password")
	}

	res := db.Create(&usersRepo.Users{
		ID:       uuid.NewV4().String(),
		Name:     "Admin",
		Email:    "admin@gmail.com",
		Password: hashedPassword,
		IsAdmin:  true,
	})
	if res.Error != nil {
		fmt.Println("Failed creating first two data")
	}

	res = db.Create(&usersRepo.Users{
		ID:       uuid.NewV4().String(),
		Name:     "User",
		Email:    "user@gmail.com",
		Password: hashedPassword,
		IsAdmin:  false,
	})
	if res.Error != nil {
		fmt.Println("Failed creating first two data")
	}
}

func main() {
	configDb := mysql.ConfigDB{
		DB_URL:      os.Getenv("DB_URL"),
		DB_Username: os.Getenv("DB_USER"),
		DB_Password: os.Getenv("DB_PASSWORD"),
		DB_Host:     os.Getenv("DB_HOST"),
		DB_Port:     os.Getenv("DB_PORT"),
		DB_Database: os.Getenv("DB_NAME"),
	}

	configJWT := middlewares.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	var db *gorm.DB
	if configDb.DB_Host == "" && configDb.DB_Port == "" {
		db = configDb.InitialDBWithURL()
	} else {
		db = configDb.InitialDB()
	}

	dbMigrate(db)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	e := echo.New()

	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"http://localhost:3000"},
	// 	AllowMethods: []string{
	// 		http.MethodGet,
	// 		http.MethodHead,
	// 		http.MethodPut,
	// 		http.MethodPatch,
	// 		http.MethodPost,
	// 		http.MethodDelete},
	// }))

	// users
	usersRepoInterface := usersRepo.NewUsersRepository(db)
	usersUseCaseInterface := usersUseCase.NewUsersUseCase(usersRepoInterface, timeoutContext, &configJWT)
	usersUseControllerInterface := usersController.NewUsersController(usersUseCaseInterface)

	routesInit := routes.RouteControllerList{
		UsersController: *usersUseControllerInterface,
		JWTConfig:       configJWT.Init(),
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
