package models

import "gorm.io/gorm"

type Ingredient struct {
	gorm.Model
	ID        uint
	Name      string `json:"name"`
	Amount 	  string `json:"amount"`
	RecipeID  uint // Foreign key to Recipe
	Recipe    Recipe
}