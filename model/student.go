package model

import (
	"gorm.io/gorm"
	"time"
)

type Student struct {
	gorm.Model
	Name   string
	Number string
	Gender int
	Birth  time.Time
}
