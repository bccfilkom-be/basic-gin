package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
    // Membuat Gin Engine
    r := gin.Default()

    // Membuat route "/helloworld"
    r.GET("/helloworld", func(c *gin.Context) {
        // Mengirimkan string "hello world" sebagai response
        c.String(200, "hello world")
    })

    // Menjalankan Gin Engine
    r.Run("localhost:8080")
}