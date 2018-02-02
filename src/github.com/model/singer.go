package model

import "github.com/jinzhu/gorm"

type Singer struct {
	Name string
	Meta string
	Active string
	Avatar string
	gorm.InitModel
}

