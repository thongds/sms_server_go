package model

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	Name string
	Active string
	gorm.InitModel
}
