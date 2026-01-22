package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/moraiba/karaoke-queue/internal/db"
)

func main() {

	// Puerto (listo para Docker / prod)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	database := db.Connect()
	defer database.Close()

	// Migraciones
	db.RunMigrations(database)

	// Gin router
	router := gin.Default()

	// Health check (OBLIGATORIO)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	log.Printf("🚀 Server running on :%s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
