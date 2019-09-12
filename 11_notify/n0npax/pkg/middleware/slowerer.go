package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

// SlowRequest introduces delay
func SlowRequest(duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		wait := duration - time.Since(start)
		// negative duration is treated as 0
		time.Sleep(wait)
	}
}
