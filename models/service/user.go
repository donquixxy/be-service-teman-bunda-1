package service

import "time"

type User struct {
	Id           string
	Username     string
	Password     string
	CreatedDate  time.Time
	RefreshToken string
}
