package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"recipe-app/handlers"
)

const (
	dbConnStr = "host=postgres port=5432 user=postgres password=password dbname=recipe_app sslmode=disable"
)


func main() {
	// Connect to PostgreSQL
	db, err := sqlx.Connect("postgres", dbConnStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	defer db.Close()

	// Initialize database schema
	handlers.InitDB(db)

	// Set up Gin router
	r := gin.Default()

	// Serve static files
	r.Static("/static", "./static")

	// Serve HTML templates
	r.LoadHTMLGlob("static/*.html")

	// Routes

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// API endpoints
	api := r.Group("/api")
	{
		api.GET("/recipes", handlers.GetRecipes(db))
		api.POST("/recipes", handlers.CreateRecipe(db))
		api.GET("/menu", handlers.GetWeeklyMenu(db))
		api.POST("/menu", handlers.CreateWeeklyMenu(db))
		api.GET("/shopping-list", handlers.GetShoppingList(db))
	}

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
