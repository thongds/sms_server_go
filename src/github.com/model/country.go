package model

import (
	"github.com/jinzhu/gorm"
)

type Country struct {
	Name string
	Active string

	gorm.InitModel
}


func (ctr *Country)AddNewCountry(){
	if cnn == nil {
		cnn = ConnectDB()
	}
	cnn.Create(&ctr)
}
func (ctr *Country)GetCountry() *[]Country{
	if cnn == nil {
		cnn = ConnectDB()
	}
	country := &[]Country{}
	cnn.Find(country)
	return country
}
func (ctr *Country)GetCountryById(id string) *Country{
	if cnn == nil {
		cnn = ConnectDB()
	}
	country := &Country{}
	cnn.Where("id = ?",id).Find(country)
	return country
}