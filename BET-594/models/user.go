package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID                    uint
	Firstname             string `json:"firstname"`
	Lastname 	          string `json:"lastname"`
	Email    	          string `json:"email" gorm:"unique"`
	Password 	          string `json:"password"`
	Recipes 	          []Recipe
	FavoriteRecipes  	  []Recipe `gorm:"many2many:user_favorite_recipes;constraint:OnDelete:CASCADE;"`
}


func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
