package model

import "gorm.io/gorm"

type MergeUsersLanguages struct {
	gorm.Model
	UserID     int `gorm:"primaryKey"`
	User       *User
	LanguageID int `gorm:"primaryKey"`
	Language   *Language
}
