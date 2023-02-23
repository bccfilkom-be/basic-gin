package main

import (
	"basic-gin/database"
	"basic-gin/handler"
	middleware "basic-gin/middleware"
	"basic-gin/repository"
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

	postRepo := repository.NewPostRepository(db)
	commentRepo := repository.NewCommentRepository(db)
	userRepo := repository.NewUserRepository(db)

	// HANDLERS
	postHandler := handler.NewPostHandler(&postRepo)
	commentHandler := handler.NewCommentHandler(&commentRepo)
	userHandler := handler.NewUserHandler(&userRepo)

	// ROUTES
	r.GET("/helloworld", func(c *gin.Context) {
		// Mengirimkan string "hello world" sebagai response
		c.String(200, "hello world")
	})

	r.POST("/user/register", userHandler.CreateUser)
	r.POST("/user/login", userHandler.LoginUser)
	// Untuk menambahkan middleware, tambahkan pada parameter kedua middleware nya, diikuti handler pada parameter ketiga
	r.GET("/user/:id", userHandler.GetUserById)

	r.POST("/post", middleware.JwtMiddleware(), postHandler.CreatePost)
	r.GET("/post/:id", postHandler.GetPostByID)
	r.GET("/posts", postHandler.GetAllPost)
	r.PATCH("/post/:id", postHandler.UpdatePostByID)
	r.DELETE("/post/:id", postHandler.DeletePostByID)

	r.POST("/comment", commentHandler.CreateNewComment)
	r.GET("/comment/:id", commentHandler.GetCommentByID)
	r.GET("/comment", commentHandler.GetCommentByTitleQuery)
	r.PATCH("/comment/:id", commentHandler.UpdateCommentByID)
	r.DELETE("/comment/:id", commentHandler.DeleteCommentByID)

	// Menjalankan Gin Engine
	r.Run(":" + port)
}
