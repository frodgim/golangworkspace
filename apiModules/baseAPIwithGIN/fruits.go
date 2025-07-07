package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Fruit structure
type Fruit struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// In-memory fruit store
var fruits = []Fruit{}

func RegisterFruitsAPI(router *gin.Engine) {
	router.GET("/fruits", getFruits)
	router.POST("/fruits", postFruit)
	router.PUT("/fruits/:name", putFruit)
	router.DELETE("/fruits/:name", deleteFruit)
}

// Validation helper
func validateFruit(f Fruit) (bool, string) {
	if strings.TrimSpace(f.Name) == "" {
		return false, "Name must not be empty"
	}
	if len(f.Name) > 50 {
		return false, "Name must be at most 50 characters"
	}
	if f.Age <= 0 {
		return false, "Age must be greater than zero"
	}
	return true, ""
}

func getFruits(c *gin.Context) {
	c.JSON(http.StatusOK, fruits)
}

func postFruit(c *gin.Context) {
	var newFruit Fruit
	if err := c.ShouldBindJSON(&newFruit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if ok, msg := validateFruit(newFruit); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}
	fruits = append(fruits, newFruit)
	c.JSON(http.StatusCreated, newFruit)
}

func putFruit(c *gin.Context) {
	name := c.Param("name")
	var updatedFruit Fruit
	if err := c.ShouldBindJSON(&updatedFruit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if ok, msg := validateFruit(updatedFruit); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}
	for i, f := range fruits {
		if f.Name == name {
			fruits[i] = updatedFruit
			c.JSON(http.StatusOK, updatedFruit)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Fruit not found"})
}

func deleteFruit(c *gin.Context) {
	name := c.Param("name")
	for i, f := range fruits {
		if f.Name == name {
			fruits = append(fruits[:i], fruits[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Fruit deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Fruit not found"})
}
