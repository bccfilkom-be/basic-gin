package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("failed to load env file")
	}
	port := os.Getenv("PORT")

	// Membuat Gin Engine
    r := gin.Default()

    // Membuat route "/helloworld"
    r.GET("/helloworld", func(c *gin.Context) {
        // Mengirimkan string "hello world" sebagai response
        c.String(200, "hello world")
    })

    // Menjalankan Gin Engine
    r.Run(":" + port)
}