package service

import "time"

type User struct {
	Id           string
	Username     string
	Password     string
	IdKelurahan  int
	CreatedDate  time.Time
	RefreshToken string
}
