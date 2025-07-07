package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	RegisterFruitsAPI(r)
	return r
}

func TestPostFruit_Valid(t *testing.T) {
	router := setupRouter()
	fruit := Fruit{Name: "Apple", Age: 2}
	body, _ := json.Marshal(fruit)

	req, _ := http.NewRequest("POST", "/fruits", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var resp Fruit
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, fruit, resp)
}

func TestPostFruit_InvalidName(t *testing.T) {
	router := setupRouter()
	fruit := Fruit{Name: "", Age: 2}
	body, _ := json.Marshal(fruit)

	req, _ := http.NewRequest("POST", "/fruits", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPostFruit_InvalidAge(t *testing.T) {
	router := setupRouter()
	fruit := Fruit{Name: "Banana", Age: 0}
	body, _ := json.Marshal(fruit)

	req, _ := http.NewRequest("POST", "/fruits", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPostFruit_NameTooLong(t *testing.T) {
	router := setupRouter()
	longName := ""
	for i := 0; i < 51; i++ {
		longName += "a"
	}
	fruit := Fruit{Name: longName, Age: 2}
	body, _ := json.Marshal(fruit)

	req, _ := http.NewRequest("POST", "/fruits", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
