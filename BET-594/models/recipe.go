package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	ID              uint
	Name     		string `json:"name"`
	Details 	  	string `json:"details"`
	Nbr_likes 	  	int `json:"nbr_likes"`
	Ingredients 		[]Ingredient `gorm:"constraint:OnDelete:CASCADE" json:"ingredients"`
	CategoryID  uint // Foreign key to Recipe
	Category 		Category `gorm:"foreignKey:CategoryID"`
	UserID  	uint // Foreign key to Recipe
	User 		User `gorm:"foreignKey:UserID"`
	Users    		[]User `gorm:"many2many:user_favorite_recipes;"`
}