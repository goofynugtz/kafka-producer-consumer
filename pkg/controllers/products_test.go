package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/goofynugtz/kafka-producer-consumer/pkg/routes"
)

func setupRouter() *gin.Engine {
	router := gin.New()
	public := router.Group("/")
	public.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	routes.PublicRoutes(public)
	return router
}

func TestRecieveProduct(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()

	payload := `{"user_id": 100, "product_name":"Sony XG14","product_description": "Protable Speakers", "product_images": ["https://source.unsplash.com/random"], "product_price": 13640 }
	`

	req, _ := http.NewRequest("POST", "/recieve", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, w.Code)
	}
}
