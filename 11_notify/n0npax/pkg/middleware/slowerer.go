package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// SlowRequest introduces delay n miliseconds
func SlowRequest(duration string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if minDuration, err := time.ParseDuration(duration); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		} else {
			start := time.Now()
			c.Next()
			wait := minDuration - time.Since(start)
			time.Sleep(wait)
		}
	}
}
