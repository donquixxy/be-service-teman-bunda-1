package service

import "github.com/golang-jwt/jwt"

type Auth struct {
	Id           string `json:"id"`
	Username     string `json:"username"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenClaims struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	IdKelurahan int    `json:"id_kelurahan"`
	jwt.StandardClaims
}
