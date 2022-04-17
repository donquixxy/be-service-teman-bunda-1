package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/controllers"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
	"github.com/tensuqiuwulu/be-service-teman-bunda/repository/mysql"
	"github.com/tensuqiuwulu/be-service-teman-bunda/routes"
	"github.com/tensuqiuwulu/be-service-teman-bunda/services"
	"github.com/tensuqiuwulu/be-service-teman-bunda/utilities"
)

func main() {

	appConfig := config.GetConfig()

	mysqlDBConnection := mysql.NewDatabaseConnection(&appConfig.Database)

	// Timezone
	location, err := time.LoadLocation(appConfig.Timezone.Timezone)
	time.Local = location
	fmt.Println("location:", location, err)

	// Logger
	logrusLogger := utilities.NewLogger(appConfig.Log)

	// Validate
	validate := validator.New()

	e := echo.New()

	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      nil,
		ErrorMessage: "Request Timeout",
		Timeout:      5 * time.Second,
	}))
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = exceptions.ErrorHandler

	// Provinsi
	provinsiRepository := mysql.NewProvinsiRepository(&appConfig.Database)
	provinsiService := services.NewProvinsiService(appConfig.Webserver,
		mysqlDBConnection,
		logrusLogger,
		provinsiRepository)
	provinsiController := controllers.NewProvinsiController(appConfig.Webserver, provinsiService)
	routes.ProvinsiRoute(e, appConfig.Webserver, appConfig.Jwt, provinsiController)

	// Kabupaten
	kabupatenRepository := mysql.NewKabupatenRepository(&appConfig.Database)
	kabupatenService := services.NewKabupatenService(appConfig.Webserver,
		mysqlDBConnection,
		logrusLogger,
		kabupatenRepository)
	kabupatenController := controllers.NewKabupatenController(appConfig.Webserver, kabupatenService)
	routes.KabupatenRoute(e, appConfig.Webserver, kabupatenController)

	// Kecamatan
	kecamatanRepository := mysql.NewKecamatanRepository(&appConfig.Database)
	kecamatanService := services.NewKecamatanService(appConfig.Webserver,
		mysqlDBConnection,
		logrusLogger,
		kecamatanRepository)
	kecamatanController := controllers.NewKecamatanController(appConfig.Webserver, kecamatanService)
	routes.KecamatanRoute(e, appConfig.Webserver, kecamatanController)

	// Kelurahan
	kelurahanRepository := mysql.NewKelurahanRepository(&appConfig.Database)
	kelurahanService := services.NewKelurahanService(appConfig.Webserver,
		mysqlDBConnection,
		logrusLogger,
		kelurahanRepository)
	kelurahanController := controllers.NewKelurahanController(appConfig.Webserver, kelurahanService)
	routes.KelurahanRoute(e, appConfig.Webserver, kelurahanController)

	// Family
	familyRepository := mysql.NewFamilyRepository(&appConfig.Database)

	// Family Members
	familyMembersRepository := mysql.NewFamilyMembersRepository(&appConfig.Database)

	// Balance Point
	balancePointRepository := mysql.NewBalancePointRepository(&appConfig.Database)
	balancePointService := services.NewBalancePointService(appConfig.Webserver,
		mysqlDBConnection,
		logrusLogger,
		balancePointRepository)
	balancePointController := controllers.NewBalancePointController(appConfig.Webserver, logrusLogger, balancePointService)
	routes.BalancePointRoute(e, appConfig.Webserver, appConfig.Jwt, balancePointController)

	// Balance Point Tx
	balancePointTxRepository := mysql.NewBalancePointTxRepository(&appConfig.Database)

	// User
	userRepository := mysql.NewUserRepository(&appConfig.Database)
	userService := services.NewUserService(appConfig.Webserver,
		mysqlDBConnection,
		appConfig.Jwt,
		validate,
		logrusLogger,
		userRepository,
		provinsiRepository,
		familyRepository,
		familyMembersRepository,
		balancePointRepository,
		balancePointTxRepository)
	userController := controllers.NewUserController(appConfig.Webserver, logrusLogger, userService)
	routes.UserRoute(e, appConfig.Webserver, appConfig.Jwt, userController)

	// Careful shutdown
	go func() {
		if err := e.Start(":" + strconv.Itoa(int(appConfig.Webserver.Port))); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// mysql database
	mysql.Close(mysqlDBConnection)
	fmt.Println("Echo was successful shutdown.")
}
