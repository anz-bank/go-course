package puppy

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	//LostPuppyURL contains lost puppy service endpoint and it setup by cmd/puppy-server
	LostPuppyURL = ""
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
		err = s.DeletePuppy(id)
		if err != nil {
			checkError(err, c)
			return
		}
		c.JSON(http.StatusNoContent, nil)
		go LostPuppyReq(id)
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
	code := c.Writer.Status()
	// default err
	if code < 400 {
		code = ErrCodeInternal
	}
	message := err.Error()
	// puppy error
	if myErr, ok := err.(*Error); ok {
		code = myErr.Code
		message = myErr.Message
	}
	c.JSON(code, gin.H{"message": message})
}
