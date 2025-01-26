// api/handlers_test.go
package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/memory"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	userHandler := UserHandler{DB: db}
	r.POST("/users", userHandler.CreateUser)
	return r
}

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin_TEST_MODE)
	db, _ := gorm.Open(memory.New(), &gorm.Config{})
	db.AutoMigrate(&User{})

	router := setupRouter(db)

	// Test creating a user
	user := User{Firstname: "John", Lastname: "Doe", Email: "johndoe@example.com", Age: 30}
	body, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusCreated {
		t.Fatalf("Expected status code 201, got %d", res.Code)
	}
}
