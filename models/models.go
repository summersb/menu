package models

type Recipe struct {
	ID          int      `json:"id" db:"id"`
	Name        string   `json:"name" db:"name"`
	Ingredients []string `json:"ingredients" db:"ingredients"`
	Instructions string  `json:"instructions" db:"instructions"`
}

type WeeklyMenu struct {
	ID         int    `json:"id" db:"id"`
	DayOfWeek  string `json:"day_of_week" db:"day_of_week"`
	RecipeID   int    `json:"recipe_id" db:"recipe_id"`
	RecipeName string `json:"recipe_name" db:"name"`
}
