package main

import (
	"basic-gin/database"
	"basic-gin/handler"
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

	// Initialize database connection
	db := database.InitDB()

	// Auto migrate entities 
	if err := database.AutoMigrate(db); err != nil {
		log.Fatalln("auto migrate error,", err)
	}

	// Membuat Gin Engine
	r := gin.Default()
	
	// HANDLERS
	postHandler := handler.NewPostHandler(db)

	// ROUTES
	r.GET("/helloworld", func(c *gin.Context) {
		// Mengirimkan string "hello world" sebagai response
		c.String(200, "hello world")
	})
	r.POST("/create-post", postHandler.CreatePost)
	r.GET("/get-post/:id", postHandler.GetPostByID)
	r.GET("/posts", postHandler.GetAllPost)
	r.PATCH("/post/:id", postHandler.UpdatePostByID)

	// Menjalankan Gin Engine
	r.Run(":" + port)
}