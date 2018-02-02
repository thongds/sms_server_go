package model

import "github.com/jinzhu/gorm"

type Songs struct {
	Name string
	Active string
	CategoriesId uint
	SingersId uint
	CountriesId uint
	SourceAudio string
	SourceLyric string
	Meta string
	Image string
	Thumb string
	gorm.InitModel
}

