package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/controllers"
)

// Provinsi Route
func ProvinsiRoute(e *echo.Echo, configWebserver config.Webserver, configJWT config.Jwt, provinsiControllerInterface controllers.ProvinsiControllerInterface) {
	group := e.Group("api/v1")
	group.GET("/provinsi", provinsiControllerInterface.FindAllProvinsi)
}

// Kabupaten Route
func KabupatenRoute(e *echo.Echo, configWebserver config.Webserver, kabupatenControllerInterface controllers.KabupatenControllerInterface) {
	group := e.Group("api/v1")
	group.GET("/kabupaten/provinsi/:id", kabupatenControllerInterface.FindAllKabupatenByIdProvinsi)
}

// Kabupaten Route
func KecamatanRoute(e *echo.Echo, configWebserver config.Webserver, kecamatanControllerInterface controllers.KecamatanControllerInterface) {
	group := e.Group("api/v1")
	group.GET("/kecamatan/kabupaten/:id", kecamatanControllerInterface.FindAllKecamatanByIdKabupaten)
}

// Kelurahan Route
func KelurahanRoute(e *echo.Echo, configWebserver config.Webserver, kelurahanControllerInterface controllers.KelurahanControllerInterface) {
	group := e.Group("api/v1")
	group.GET("/kelurahan/kecamatan/:id", kelurahanControllerInterface.FindAllKelurahanByIdKecamatan)
}

// User Route
func UserRoute(e *echo.Echo, configWebserver config.Webserver, configurationJWT config.Jwt, userControllerInterface controllers.UserControllerInterface) {
	group := e.Group("api/v1")
	group.POST("/user/create", userControllerInterface.CreateUser)
	group.GET("/user/referal/:referal", userControllerInterface.FindUserByReferal)
	group.GET("/user/:id", userControllerInterface.FindUserById)
}

// User Route
func BalancePointRoute(e *echo.Echo, configWebserver config.Webserver, configurationJWT config.Jwt, balancePointControllerInterface controllers.BalancePointControllerInterface) {
	group := e.Group("api/v1")
	group.GET("/balance_point_with_tx/user/:id", balancePointControllerInterface.FindBalancePointWithTxByIdUser)
	group.GET("/balance_point/user/:id", balancePointControllerInterface.FindBalancePointWithTxByIdUser)
}
