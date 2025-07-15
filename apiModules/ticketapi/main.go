package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

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

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}
	r.Run(fmt.Sprintf(":%s", appPort)) // listen and serve on ":8080" or the port specified in APP_PORT)
}

func cacheTicket(id string, ticket Ticket) {
	logrus.Debugf("Caching ticket %s", id)
	jsonData, _ := json.Marshal(ticket)
	rdb.Set(ctx, "ticket:"+id, jsonData, 0)
	// Increment frequency
	rdb.ZIncrBy(ctx, "ticket:freq", 1, id)
	// Keep only top 50
	rdb.ZRemRangeByRank(ctx, "ticket:freq", 0, -51)
	// Remove cache for IDs not in top 50
	topIDs, _ := rdb.ZRevRange(ctx, "ticket:freq", 0, 49).Result()
	allIDs, _ := rdb.ZRange(ctx, "ticket:freq", 0, -1).Result()
	idSet := make(map[string]struct{})
	for _, tid := range topIDs {
		idSet[tid] = struct{}{}
	}
	for _, tid := range allIDs {
		if _, ok := idSet[tid]; !ok {
			rdb.Del(ctx, "ticket:"+tid)
		}
	}
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
	cacheTicket(strconv.Itoa(int(ticket.ID)), ticket)
	c.JSON(http.StatusCreated, ticket)
}

func getTicket(c *gin.Context) {
	id := c.Param("id")
	var ticket Ticket
	if val, err := rdb.Get(ctx, "ticket:"+id).Result(); err == nil {
		logrus.Debugf("Cache hit for ticket %s", id)
		json.Unmarshal([]byte(val), &ticket)
		// Increment frequency and maintain cache
		rdb.ZIncrBy(ctx, "ticket:freq", 1, id)
		c.JSON(http.StatusOK, ticket)
		return
	}
	if err := db.First(&ticket, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}
	cacheTicket(id, ticket)
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
	// Refresh cache and frequency
	cacheTicket(id, ticket)
	c.JSON(http.StatusOK, ticket)
}

func deleteTicket(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&Ticket{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}
	logrus.Debugf("Deleting ticket %s from cache", id)
	rdb.Del(ctx, "ticket:"+id)
	c.Status(http.StatusNoContent)
}

func listTickets(c *gin.Context) {
	var tickets []Ticket
	db.Find(&tickets)
	c.JSON(http.StatusOK, tickets)
}
