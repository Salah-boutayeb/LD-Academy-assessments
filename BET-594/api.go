package main

import (
	"api/controllers"
	"api/database"
	"api/middlewares"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)
func main() {
	// Initialize Database
	godotenv.Load(".env")
	urldb := "postgres://"+os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@"+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+"/"+os.Getenv("DB_NAME")
	
	
	// database.Connect("postgres://postgres:salah2000@localhost:5432/go_api")
	database.Connect(urldb)

	// database.Connect("root:root@tcp(localhost:3306)/go_api?parseTime=true")
	database.Migrate()
	
	// Initialize Router
	router := initRouter()
	router.Run(":8080")
}
func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/login", controllers.GenerateToken)
		api.POST("/test", controllers.TestContoller)
		api.POST("/user/register", controllers.RegisterUser)
		
		recipeRoutes := api.Group("/recipes").Use(middlewares.Auth())
		{
			recipeRoutes.GET("/", controllers.GetAllRecipesWithIngredients)
			recipeRoutes.GET("/:id", controllers.GetRecipeWithIngredientsByID)
			recipeRoutes.POST("/", controllers.CreateRecipeWithIngredients)
			recipeRoutes.DELETE("/:id", controllers.DeleteRecipeByID)
			recipeRoutes.PUT("/:id", controllers.UpdateRecipe)
			recipeRoutes.GET("/category/:id", controllers.GetRecipesByCategory)
			// *******
		}
		categoryRoutes := api.Group("/categories").Use(middlewares.Auth())
		{
			categoryRoutes.POST("/", controllers.CreateCategory)
			categoryRoutes.GET("/", controllers.GetAllCategoriesWithChildren)
			categoryRoutes.GET("/:id", controllers.GetCategoryWithChildrenByID)
			categoryRoutes.DELETE("/:id", controllers.DeleteCategoryByID)
			categoryRoutes.PUT("/:id", controllers.UpdateCategory)
		}
		userRoutes := api.Group("/users").Use(middlewares.Auth())
		{
			userRoutes.GET("/:id/recipes", controllers.GetRecipesByUserID)
			userRoutes.POST("/:id/favorite_recipes/:recipeId", controllers.AddRecipeToFavorites)
			userRoutes.GET("/:id/favorite_recipes", controllers.GetUserfavoriteRecipes)
		}
	}
	return router
}