package models

import "github.com/jinzhu/gorm"

var db *gorm.DB

func init() {}

func GetDB() *gorm.DB {
	return db
}
