package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Ticket struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"size:50;not null"`
	Type string `json:"type" gorm:"size:5;not null"`
}

var (
	db  *gorm.DB
	rdb *redis.Client
	ctx = context.Background()
)

const isMySQLLocal bool = true

var validTypes = map[string]bool{"kindA": true, "kindB": true, "kindC": true}

func main() {
	if isMySQLLocal {
		_ = godotenv.Load(".env")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "3306"
	}
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + dbPort + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = dbConn
	db.AutoMigrate(&Ticket{})

	rdb = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	r := gin.Default()

	r.POST("/tickets", createTicket)
	r.GET("/tickets/:id", getTicket)
	r.PUT("/tickets/:id", updateTicket)
	r.DELETE("/tickets/:id", deleteTicket)
	r.GET("/tickets", listTickets)

	r.Run(":8080")
}

func createTicket(c *gin.Context) {
	var ticket Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(ticket.Name) > 50 || !validTypes[ticket.Type] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid name or type"})
		return
	}
	db.Create(&ticket)
	c.JSON(http.StatusCreated, ticket)
}

func getTicket(c *gin.Context) {
	id := c.Param("id")
	var ticket Ticket
	if val, err := rdb.Get(ctx, "ticket:"+id).Result(); err == nil {
		c.Data(http.StatusOK, "application/json", []byte(val))
		return
	}
	if err := db.First(&ticket, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}
	c.JSON(http.StatusOK, ticket)
}

func updateTicket(c *gin.Context) {
	id := c.Param("id")
	var ticket Ticket
	if err := db.First(&ticket, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}
	var input Ticket
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(input.Name) > 50 || !validTypes[input.Type] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid name or type"})
		return
	}
	db.Model(&ticket).Updates(input)
	rdb.Del(ctx, "ticket:"+id)
	c.JSON(http.StatusOK, ticket)
}

func deleteTicket(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&Ticket{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}
	rdb.Del(ctx, "ticket:"+id)
	c.Status(http.StatusNoContent)
}

func listTickets(c *gin.Context) {
	var tickets []Ticket
	db.Find(&tickets)
	c.JSON(http.StatusOK, tickets)
}
