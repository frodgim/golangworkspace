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

type testTicket struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/tickets", createTicket)
	r.GET("/tickets/:id", getTicket)
	r.PUT("/tickets/:id", updateTicket)
	r.DELETE("/tickets/:id", deleteTicket)
	r.GET("/tickets", listTickets)
	return r
}

func TestCreateTicket(t *testing.T) {
	r := setupRouter()
	ticket := testTicket{Name: "Test Ticket", Type: "kindA"}
	jsonValue, _ := json.Marshal(ticket)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/tickets", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
}

func TestCreateTicketInvalidType(t *testing.T) {
	r := setupRouter()
	ticket := testTicket{Name: "Test Ticket", Type: "invalid"}
	jsonValue, _ := json.Marshal(ticket)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/tickets", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

// More tests for GET, PUT, DELETE can be added similarly.
