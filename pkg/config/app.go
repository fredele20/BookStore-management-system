package config

import "github.com/jinzhu/gorm"


var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "fredel:lyfgoes1/simplerest?charset=utf8&parseTime=True&loc=local")
	if err != nil {
		panic(err)
	}
	db = d
}


func GetDB() *gorm.DB {
	return db
}