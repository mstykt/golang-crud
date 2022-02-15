package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"golang-crud/configs"
	"golang-crud/pkg/controller"
	"golang-crud/pkg/repo"
	"golang-crud/pkg/routes"
	"golang-crud/pkg/service"
	"strconv"
)

func Run() {
	setConfig()
	routesSetup := initializeRoutes()
	runWeb(routesSetup)
}

func runWeb(routes *routes.Routes) {
	e := echo.New()
	routes.InitRoutes(e)
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(configs.Conf.Server.Port)))
}

func setConfig() {
	viper.SetConfigName("application")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configs.Conf)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	dbErr := viper.Unmarshal(&configs.DBConf)
	if dbErr != nil {
		panic(dbErr)
	}
}

func initializeRoutes() *routes.Routes {
	connection := configs.NewConnection()
	userRepo := repo.NewUserRepo(connection)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	return routes.NewRoutes(userController)
}
