package puppy

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

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
		c.JSON(200, p)
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
		c.JSON(204, p)
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
		c.JSON(201, gin.H{"id": id})
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
		c.JSON(204, "")
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
