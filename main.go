package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mezink-test/src/handlers"
	"github.com/mezink-test/src/services/db"
)

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func route(router *gin.Engine) {
	api := router.Group("/api/v1")

	api.GET("/records", handlers.GetRecordsHandler)
	api.POST("/records", handlers.CreateRecordHandler)
}

func main() {
	var err error
	gin.SetMode(os.Getenv("GIN_MODE"))

	// show panic when db not connect
	_, err = db.InitDB()
	if err != nil {
		panic("Failed to connect to database")
	}
	log.Printf("Database successfully start")

	// set rest route
	router := gin.Default()
	route(router)

	port := os.Getenv("API_PORT")
	log.Printf("Starting user service. PORT: %s\n", port)
	router.Run(":" + port)
}
