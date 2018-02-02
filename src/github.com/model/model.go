package model


import (
	"github.com/database"
	"github.com/jinzhu/gorm"
)

var db database.Mysql
var cnn *gorm.DB
func ConnectDB() *gorm.DB  {
	cnn = db.New()
	return cnn
}
func GetDataById(data interface{},id string) interface{}{
	if cnn == nil {
		cnn = ConnectDB()
	}
	cnn.Where("id = ?",id).Find(data)
	return data
}
func GetDataByField(data interface{},field string,value string) interface{} {
	if cnn == nil {
		cnn = ConnectDB()
	}
	cnn.Where(field+" = ?",value).Find(data)

	return data
}
func GetAllData(table interface{}){
	if cnn == nil {
		cnn = ConnectDB()
	}
	cnn.Find(table)
}
func InsertNewData(record interface{})  {
	if cnn == nil {
		cnn = ConnectDB()
	}
	cnn.Create(record)
}