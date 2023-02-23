package response

import "github.com/gin-gonic/gin"

// bentuk response nya terserah lah yang penting konsisten 

func Success(c *gin.Context, httpCode int, msg string, data interface{}) {
	switch httpCode / 100 {
		case 2:
			c.JSON(httpCode, map[string]interface{}{
				"status":  "success",
				"message": msg,
				"data":    data,
			})
		default:
			c.JSON(500, map[string]interface{}{
				"status":  "error",
				"message": "RESPONSE ERROR",
			})
	} 
}

func FailOrError(c *gin.Context, httpCode int, msg string, err error) {
	switch httpCode / 100 {
		case 4: //FAIL 4xx
			c.JSON(httpCode, gin.H{
				"status":  "fail",
				"message": msg,
				"data":  gin.H{
					"error" : err.Error(),
				},
			})
		case 5: //ERROR 5xx
			c.JSON(httpCode, gin.H{
				"status":  "error",
				"message": msg,
			})
			
		default:
			c.JSON(500, gin.H{
				"status":  "error",
				"message": "RESPONSE ERROR",
			})
	}
}
