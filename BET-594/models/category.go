package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID           uint
	Name         string `json:"name"`
	ParentID     uint `gorm:"default:null"`
	Categories     []Category `gorm:"foreignKey:ParentID;constraint:OnDelete:CASCADE;"`

	Recipes 	[]Recipe
}


