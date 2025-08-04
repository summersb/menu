package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func TestGetRecipes(t *testing.T) {
    db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres password=password dbname=recipe_app sslmode=disable")
    if err != nil {
        t.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    InitDB(db)

    router := gin.Default()
    router.GET("/api/recipes", GetRecipes(db))

    req, _ := http.NewRequest("GET", "/api/recipes", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Errorf("Expected status 200, got %d", w.Code)
    }
}