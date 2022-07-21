package model

import "gorm.io/gorm"

type CreditCard struct {
	gorm.Model
	Number string
	UserId string
	User   *User
}
