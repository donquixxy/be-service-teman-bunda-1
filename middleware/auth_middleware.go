package middleware

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	modelService "github.com/tensuqiuwulu/be-service-teman-bunda/models/service"
)

func Authentication(configurationJWT config.Jwt) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:       &modelService.TokenClaims{},
		SigningKey:   []byte(configurationJWT.Key),
		ErrorHandler: ErrorHandler,
	})
}

func ErrorHandler(err error) error {
	fmt.Println(err)
	return err
}

func TokenClaimsIdUser(c echo.Context) (id string) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*modelService.TokenClaims)
	idUser := claims.Id
	return idUser
}

func TokenClaimsIdKelurahan(c echo.Context) (id int) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*modelService.TokenClaims)
	idKelurahan := claims.IdKelurahan
	return idKelurahan
}
