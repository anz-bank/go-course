package puppy

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"

	middleware "github.com/anz-bank/go-course/11_notify/n0npax/pkg/middleware"
	"github.com/gin-gonic/gin"
)

var (
	LostPuppyURL        = ""
	SlowRequestDuration = "2s"
	PuppyDeleteNotifyF  = LostPuppyReq
)

// LostPuppyReq sends request to lostpuppy service
func LostPuppyReq(id int) {
	jsonBytes := []byte(fmt.Sprintf(`{"id":%d}`, id))
	resp, err := http.Post(LostPuppyURL, "application/json", bytes.NewBuffer(jsonBytes)) // #nosec
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	log.Printf("lostpuppy service response code: %d (puppy id: %d)", resp.StatusCode, id)
}

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

// RestBackend serves puppy over http
func RestBackend(s Storer) *gin.Engine {
	r := gin.Default()

	r.GET("/api/puppy/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			checkError(err, c)
			return
		}
		p, err := s.ReadPuppy(id)
		if err != nil {
			checkError(err, c)
			return
		}
		c.JSON(http.StatusOK, p)
	})

	r.DELETE("/api/puppy/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			checkError(err, c)
			return
		}
		p, err := s.DeletePuppy(id)
		if err != nil {
			checkError(err, c)
			return
		}
		c.JSON(http.StatusNoContent, p)
		go PuppyDeleteNotifyF(id)
	})

	r.POST("/api/puppy/", func(c *gin.Context) {
		var p Puppy
		err := c.BindJSON(&p)
		if err != nil {
			checkError(err, c)
			return
		}
		id, err := s.CreatePuppy(&p)
		if err != nil {
			checkError(err, c)
			return
		}
		c.JSON(http.StatusCreated, gin.H{"id": id})
	})

	r.PUT("/api/puppy/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			checkError(err, c)
			return
		}
		var p Puppy
		if err = c.BindJSON(&p); err != nil {
			checkError(err, c)
			return
		}
		if err = s.UpdatePuppy(id, &p); err != nil {
			checkError(err, c)
			return
		}
		c.JSON(http.StatusNoContent, "")
	})

	return r
}

func checkError(err error, c *gin.Context) {
	// default error
	code := ErrCodeInternal
	message := err.Error()
	// puppy error
	if myErr, ok := err.(*Error); ok {
		code = myErr.Code
		message = myErr.Message
	}
	c.JSON(code, gin.H{"message": message})
}
