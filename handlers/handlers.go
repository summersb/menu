package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"recipe-app/models"
)


// InitDB creates the necessary database tables

func InitDB(db *sqlx.DB) {
	schema := `
	CREATE TABLE IF NOT EXISTS recipes (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		ingredients JSONB NOT NULL,
		instructions TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS weekly_menu (
		id SERIAL PRIMARY KEY,
		day_of_week TEXT NOT NULL,
		recipe_id INTEGER REFERENCES recipes(id)
	);`
	_, err := db.Exec(schema)
	if err != nil {
		panic("Failed to initialize database: " + err.Error())

	}
}

// GetRecipes returns all recipes
func GetRecipes(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var recipes []models.Recipe
		err := db.Select(&recipes, "SELECT id, name, ingredients, instructions FROM recipes")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, recipes)
	}
}

// CreateRecipe adds a new recipe
func CreateRecipe(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var recipe models.Recipe
		if err := c.ShouldBindJSON(&recipe); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := db.QueryRow(
			"INSERT INTO recipes (name, ingredients, instructions) VALUES ($1, $2, $3) RETURNING id",
			recipe.Name, recipe.Ingredients, recipe.Instructions,
		).Scan(&recipe.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, recipe)
	}
}

// GetWeeklyMenu returns the weekly menu
func GetWeeklyMenu(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var menu []models.WeeklyMenu
		err := db.Select(&menu, `
			SELECT wm.id, wm.day_of_week, wm.recipe_id, r.name
			FROM weekly_menu wm

			JOIN recipes r ON wm.recipe_id = r.id
			ORDER BY wm.id`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, menu)
	}
}

// CreateWeeklyMenu adds a recipe to the weekly menu
func CreateWeeklyMenu(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var menu models.WeeklyMenu
		if err := c.ShouldBindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := db.QueryRow(

			"INSERT INTO weekly_menu (day_of_week, recipe_id) VALUES ($1, $2) RETURNING id",
			menu.DayOfWeek, menu.RecipeID,
		).Scan(&menu.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, menu)
	}
}

// GetShoppingList generates a shopping list from the weekly menu

func GetShoppingList(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var recipes []models.Recipe

		err := db.Select(&recipes, `
			SELECT r.ingredients
			FROM weekly_menu wm
			JOIN recipes r ON wm.recipe_id = r.id`)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ingredients := make(map[string]bool)
		for _, recipe := range recipes {
			for _, ing := range recipe.Ingredients {
				ingredients[ing] = true
			}
		}

		var shoppingList []string
		for ing := range ingredients {
			shoppingList = append(shoppingList, ing)
		}
		c.JSON(http.StatusOK, shoppingList)
	}
}
