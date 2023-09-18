package controllers

import (
	"api/database"
	"api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func addRecipe(context *gin.Context) {
	var recipe models.Recipe
	if err := context.ShouldBindJSON(&recipe); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	
	record := database.Instance.Create(&recipe)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"recipeId": recipe.ID,})
}

func GetAllRecipesWithIngredients(c *gin.Context) {
    var recipes []models.Recipe

    // Retrieve all recipes along with their associated ingredients
    if err := database.Instance.Preload("Ingredients").Preload("User").Preload("Category").Find(&recipes).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, recipes)
}

func CreateRecipeWithIngredients(c *gin.Context) {
    var recipe models.Recipe
    

    // Bind the request JSON to the Recipe struct
    if err := c.ShouldBindJSON(&recipe); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error 1": err.Error()})
        return
    }
    
    // Create the recipe in the database
    if err := database.Instance.Create(&recipe).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error 2": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "Recipe created successfully", "recipe": recipe})
    
}


func DeleteRecipeByID(c *gin.Context) {
    // Parse the recipe ID from the URL parameter
    recipeIDStr := c.Param("id")
    recipeID, err := strconv.ParseUint(recipeIDStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
        return
    }

    // Check if the recipe exists
	var recipe models.Recipe
	
    
    if err := database.Instance.Where("id = ?", recipeID).Find(&recipe).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "recipe not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        return
    }
    
    // Delete the recipe
    if err := database.Instance.Where("id = ?", recipeID).Delete(&recipe).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete recipe"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "recipe deleted successfully","recipe":recipe})
}

func UpdateRecipe(c *gin.Context) {
    // Parse the recipe ID from the URL parameter
    recipeIDStr := c.Param("id")
    recipeID, err := strconv.ParseUint(recipeIDStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
        return
    }

    // Check if the recipe exists
    var recipe models.Recipe
    if err := database.Instance.Where("id = ?", recipeID).Find(&recipe).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "recipe not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        return
    }

    // Bind the JSON request body to the recipe struct to update its fields
    if err := c.ShouldBindJSON(&recipe); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Update the recipe in the database
    if err := database.Instance.Where("id = ?", recipeID).Save(&recipe).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update recipe"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "recipe updated successfully", "recipe": recipe})
}

func GetRecipeWithIngredientsByID(c *gin.Context) {
    // Parse the recipe ID from the URL parameter
    recipeIDStr := c.Param("id")
    recipeID, err := strconv.ParseUint(recipeIDStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
        return
    }

    // Check if the recipe exists
    var recipe models.Recipe
    if err := database.Instance.Preload("Ingredients").First(&recipe, recipeID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        return
    }

    c.JSON(http.StatusOK, recipe)
}


func GetRecipesByCategory(c *gin.Context)  {
    categoryIDStr := c.Param("id")
    categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
        return
    }

    // Check if the recipe exists
    var recipes []models.Recipe
    if err := database.Instance.Preload("Ingredients").First(&recipes, categoryID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Recipes not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        return
    }

    c.JSON(http.StatusOK, recipes)
}