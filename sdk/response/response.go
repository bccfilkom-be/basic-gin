package response

import "github.com/gin-gonic/gin"

// bentuk response nya terserah lah yang penting konsisten
type resp struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data any `json:"data"`
}  

func Success(c *gin.Context, httpCode int, msg string, data interface{}) {
	switch httpCode / 100 {
		case 2:
			c.JSON(httpCode, resp{
				Status: "success",
				Message: msg,
				Data: data,
			})
		default:
			c.JSON(500, resp{
				Status: "error",
				Message: "RESPONSE ERROR",
				Data: nil,
			})
	} 
}

func FailOrError(c *gin.Context, httpCode int, msg string, err error) {
	switch httpCode / 100 {
		case 4: //FAIL 4xx
			c.JSON(httpCode, resp{
				Status: "fail",
				Message: msg,
				Data: gin.H{
					"error" : err.Error(),
				},
			})
		case 5: //ERROR 5xx
			c.JSON(httpCode, resp{
				Status: "error",
				Message: msg,
				Data: nil,
			})
			
		default:
			c.JSON(500, resp{
				Status: "error",
				Message: "RESPONSE ERROR",
				Data: nil,
			})
	}
}
