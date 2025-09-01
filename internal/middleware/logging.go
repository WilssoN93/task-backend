package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func AddLoggingMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Incoming request: %s", c.Request.URL)
		c.Next()
	}
}
