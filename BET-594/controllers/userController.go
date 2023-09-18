package controllers

import (
	"api/database"
	"api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
func RegisterUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "firstname": user.Firstname, "lastname": user.Lastname})
}

func GetRecipesByUserID(c *gin.Context) {
    // Parse the user ID from the URL parameter
    userIDStr := c.Param("id")
    userID, err := strconv.ParseUint(userIDStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }
	var recipes []models.Recipe
    if err := database.Instance.Where("user_id = ?", userID).Preload("Ingredients").Find(&recipes).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        return
    }
    c.JSON(http.StatusOK, recipes)
}

func GetUserfavoriteRecipes(c *gin.Context)  {
	userIDStr := c.Param("id")
    userID, err := strconv.ParseUint(userIDStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }
	// Check if the user exists
    var user models.User
    if err := database.Instance.Preload("FavoriteRecipes").First(&user, userID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        return
    }

    c.JSON(http.StatusOK, user.FavoriteRecipes)
	
}



func AddRecipeToFavorites(c *gin.Context) {
    // Parse the user ID from the URL parameter
    userIDStr := c.Param("id")
    userID, err := strconv.ParseUint(userIDStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    // Parse the recipe ID from the request body or URL parameter
    recipeIDStr := c.Param("recipeId")
    recipeID, err := strconv.ParseUint(recipeIDStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
        return
    }

    // Check if the user and recipe exist
    var user models.User
    if err := database.Instance.First(&user, userID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        return
    }

    var recipe models.Recipe
    if err := database.Instance.First(&recipe, recipeID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        return
    }
	recipe.Nbr_likes ++
	database.Instance.Save(&recipe)
	

    // Add the recipe to the user's favorites
    if err := database.Instance.Model(&user).Association("FavoriteRecipes").Append(&recipe); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add recipe to favorites"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Recipe added to favorites successfully"})
}
