package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"learn-gorm/model"
	"log"
)

var db *gorm.DB
var err error

const (
	account = "root"
	pwd     = "root"
	addr    = "127.0.0.1:3306"
	dbName  = "gorm"
)

func init() {
	// 初始化数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		account, pwd, addr, dbName,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false, // 在迁移 schema 时自动生产外键(默认)
	})
	if err != nil {
		log.Fatalf("connect to mysql databse failed, err = %v", err)
		return
	}

	// 自动迁移表
	err = db.AutoMigrate(
		&model.Student{},
		&model.User{},
		&model.CreditCard{},
		&model.Language{},
		&model.MergeUsersLanguages{},
	)
	if err != nil {
		log.Fatalf("migrate schema failed, err = %v", err)
		return
	}

	log.Println("gorm init success")
}

func GetDb() *gorm.DB {
	return db
}
