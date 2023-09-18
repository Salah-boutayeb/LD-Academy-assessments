package controllers

import (
	"api/database"
	"api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Category struct {
    Name string `json:"name"`
}

// Handle POST request to create a new category
func CreateCategory(c *gin.Context) {
    var category models.Category

    // Bind the JSON request body to the Category struct
    if err := c.ShouldBindJSON(&category); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	
	
    
	database.Instance.Create(&category)
	

    // Return a success response
    c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully", "category": category})
}

func GetAllCategoriesWithChildren(c *gin.Context) {
    var categories []models.Category

    // Retrieve all top-level categories (categories without parents)
    database.Instance.Where("parent_id IS NULL").Find(&categories)

    // Recursively fetch children for each top-level category
    for i := range categories {
        categories[i].Categories = getChildren(&categories[i], database.Instance)
    }

    c.JSON(http.StatusOK, categories)
}

func getChildren(category *models.Category, db *gorm.DB) []models.Category {
    var children []models.Category
    db.Where("parent_id = ?", category.ID).Find(&children)

    // Recursively fetch children for each child category
    for i := range children {
        children[i].Categories = getChildren(&children[i], db)
    }

    return children
}


func GetCategoryWithChildrenByID(c *gin.Context) {
    // Parse the category ID from the URL parameter
    categoryIDStr := c.Param("id")
    categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
        return
    }

    // Check if the category exists
    var category models.Category
    // database.Instance.Where("id = ?", categoryID).Preload("Children").Find(&category)
    if err := database.Instance.Where("id = ?", categoryID).Preload("Recipes").Find(&category).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Category not found","err":err})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        return
    }
    category.Categories = getChildren(&category, database.Instance)

    c.JSON(http.StatusOK, category)
}
func DeleteCategoryByID(c *gin.Context) {
    // Parse the category ID from the URL parameter
    categoryIDStr := c.Param("id")
    categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
        return
    }

    // Check if the category exists
	var category Category
	
    
    if err := database.Instance.Where("id = ?", categoryID).Find(&category).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        return
    }
    
    // Delete the category
    if err := database.Instance.Where("id = ?", categoryID).Delete(&category).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully","category":category})
}

func UpdateCategory(c *gin.Context) {
    // Parse the category ID from the URL parameter
    categoryIDStr := c.Param("id")
    categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
        return
    }

    // Check if the category exists
    var category Category
    if err := database.Instance.Where("id = ?", categoryID).Find(&category).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        return
    }

    // Bind the JSON request body to the Category struct to update its fields
    if err := c.ShouldBindJSON(&category); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Update the category in the database
    if err := database.Instance.Where("id = ?", categoryID).Save(&category).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully", "category": category})
}

func GetCategoryWithRecipes(c *gin.Context)  {
    
}


