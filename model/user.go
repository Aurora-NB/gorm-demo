package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string
	CreditCard []*CreditCard `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// 使用 UserId 作为外键
	// CreditCard 中外键存储的为 User 的 ID
	// 以上两条注释皆为默认情况
	Languages []*Language `gorm:"many2many:merge_users_languages"`
}
