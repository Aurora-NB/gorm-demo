package model

import "gorm.io/gorm"

type Language struct {
	gorm.Model
	Name  string
	Users []*User `gorm:"many2many:merge_users_languages"`
}
