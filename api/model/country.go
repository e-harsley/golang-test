package model

import "gorm.io/gorm"

type Country struct {
	gorm.Model
	Name          string `gorm:"not null"`
	ShortName     string `gorm:"not null"`
	Continent     string `gorm:"not null"`
	IsOperational *bool  `gorm:"default:true"`
}
