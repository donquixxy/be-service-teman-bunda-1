package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/controllers"
	authMiddlerware "github.com/tensuqiuwulu/be-service-teman-bunda/middleware"
	modelService "github.com/tensuqiuwulu/be-service-teman-bunda/models/service"
)

var configJwt = &config.Jwt{}

var authMiddlewareJwt = middleware.JWTConfig{
	Claims:     &modelService.TokenClaims{},
	SigningKey: []byte(configJwt.Key),
}

// Provinsi Route
func ProvinsiRoute(e *echo.Echo, configWebserver config.Webserver, configJWT config.Jwt, provinsiControllerInterface controllers.ProvinsiControllerInterface) {
	group := e.Group("api/v1")
	group.GET("/provinsi", provinsiControllerInterface.FindAllProvinsi)
}

// Kabupaten Route
func KabupatenRoute(e *echo.Echo, configWebserver config.Webserver, kabupatenControllerInterface controllers.KabupatenControllerInterface) {
	group := e.Group("api/v1")
	group.GET("/kabupaten", kabupatenControllerInterface.FindAllKabupatenByIdProvinsi)
}

// Kabupaten Route
func KecamatanRoute(e *echo.Echo, configWebserver config.Webserver, kecamatanControllerInterface controllers.KecamatanControllerInterface) {
	group := e.Group("api/v1")
	group.GET("/kecamatan", kecamatanControllerInterface.FindAllKecamatanByIdKabupaten)
}

// Kelurahan Route
func KelurahanRoute(e *echo.Echo, configWebserver config.Webserver, kelurahanControllerInterface controllers.KelurahanControllerInterface) {
	group := e.Group("api/v1")
	group.GET("/kelurahan", kelurahanControllerInterface.FindAllKelurahanByIdKecamatan)
}

// User Route
func UserRoute(e *echo.Echo, configWebserver config.Webserver, configurationJWT config.Jwt, userControllerInterface controllers.UserControllerInterface) {
	group := e.Group("api/v1")
	group.POST("/user/create", userControllerInterface.CreateUser)
	group.GET("/user/referal/:referal", userControllerInterface.FindUserByReferal)
	group.GET("/user", userControllerInterface.FindUserById, authMiddlerware.Authentication(configurationJWT))
}

// Auth Route
func AuthRoute(e *echo.Echo, configWebserver config.Webserver, configurationJWT config.Jwt, authControllerInterface controllers.AuthControllerInterface) {
	group := e.Group("api/v1")
	group.POST("/auth/login", authControllerInterface.Login)
	group.POST("/auth/new-token", authControllerInterface.NewToken)
}

// Balance Point Route
func BalancePointRoute(e *echo.Echo, configWebserver config.Webserver, configurationJWT config.Jwt, balancePointControllerInterface controllers.BalancePointControllerInterface) {
	group := e.Group("api/v1")
	group.GET("/balance_point_with_tx", balancePointControllerInterface.FindBalancePointWithTxByIdUser, authMiddlerware.Authentication(configurationJWT))
	group.GET("/balance_point", balancePointControllerInterface.FindBalancePointByIdUser, authMiddlerware.Authentication(configurationJWT))
}

// Product Route
func ProductRoute(e *echo.Echo, configWebserver config.Webserver, configurationJWT config.Jwt, productControllerInterface controllers.ProductControllerInterface) {
	group := e.Group("api/v1")
	group.GET("/products", productControllerInterface.FindAllProducts, authMiddlerware.Authentication(configurationJWT))
	group.GET("/products/search", productControllerInterface.FindProductsBySearch, authMiddlerware.Authentication(configurationJWT))
	group.GET("/product/:id", productControllerInterface.FindProductById, authMiddlerware.Authentication(configurationJWT))
	group.GET("/products/category/:id", productControllerInterface.FindProductByIdCategory, authMiddlerware.Authentication(configurationJWT))
}

// Cart Route
func CartRoute(e *echo.Echo, configWebserver config.Webserver, configurationJWT config.Jwt, cartControllerInterface controllers.CartControllerInterface) {
	group := e.Group("api/v1")
	group.GET("/cart/user", cartControllerInterface.FindCartByIdUser, authMiddlerware.Authentication(configurationJWT))
}
