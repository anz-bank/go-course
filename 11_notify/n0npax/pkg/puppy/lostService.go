package puppy

import (
	"net/http"
	"time"

	middleware "github.com/anz-bank/go-course/11_notify/n0npax/pkg/middleware"
	"github.com/gin-gonic/gin"
)

var SlowRequestDuration = 2 * time.Second

// LostPuppyBackend provides information if puppy was lost
func LostPuppyBackend() *gin.Engine {
	r := gin.Default()
	r.POST("/api/lostpuppy/", middleware.SlowRequest(SlowRequestDuration), func(c *gin.Context) {
		var payload struct {
			ID int `json:"id"`
		}
		err := c.BindJSON(&payload)
		if err != nil {
			checkError(err, c)
			return
		}
		if payload.ID%2 == 0 {
			c.JSON(http.StatusCreated, gin.H{})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{})
		}
	})
	return r
}
