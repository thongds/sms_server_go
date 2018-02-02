package database

import (
	"github.com/jinzhu/gorm"
)

type Mysql struct {
	db *gorm.DB
}

func (mysql Mysql) New() *gorm.DB{
	db, _ := gorm.Open("mysql", "root:@/learnsong?charset=utf8&parseTime=True&loc=Local")
	mysql.db = db
	return mysql.db
}

