package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// JSONB is a custom type for handling JSONB fields in PostgreSQL
type JSONB []string

// Scan implements the sql.Scanner interface for JSONB
func (j *JSONB) Scan(value interface{}) error {
    bytes, ok := value.([]byte)
    if !ok {
        return errors.New("type assertion to []byte failed")
    }
    return json.Unmarshal(bytes, j)
}

// Value implements the driver.Valuer interface for JSONB
func (j JSONB) Value() (driver.Value, error) {
    if len(j) == 0 {
        return nil, nil
    }
    return json.Marshal(j)
}

type Recipe struct {
    ID          int    `json:"id" db:"id"`
    Name        string `json:"name" db:"name"`
    Ingredients JSONB  `json:"ingredients" db:"ingredients"`
    Instructions string `json:"instructions" db:"instructions"`
}

type WeeklyMenu struct {
    ID         int    `json:"id" db:"id"`
    DayOfWeek  string `json:"day_of_week" db:"day_of_week"`
    RecipeID   int    `json:"recipe_id" db:"recipe_id"`
    RecipeName string `json:"recipe_name" db:"name"`
}