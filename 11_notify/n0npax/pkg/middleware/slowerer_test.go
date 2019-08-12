package middleware

import (
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSlowRequestBadArg(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	f := SlowRequest("Invalid Input")
	f(c)
	assert.Equal(t, 500, c.Writer.Status())
}

func TestSlowRequestNegativeDuration(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	for i := -3; i >= 0; i++ {
		delay := fmt.Sprintf("%ds", i)
		expectedMaxDuration, err := time.ParseDuration("1s")
		t.Run(delay, func(t *testing.T) {
			f := SlowRequest(delay)
			start := time.Now()
			f(c)
			reqDuration := time.Since(start)
			assert.NoError(t, err)
			assert.True(t, reqDuration < expectedMaxDuration)
		})
	}
}

func TestSlowRequest(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	for i := 0; i < 3; i++ {
		delay := fmt.Sprintf("%ds", i)
		t.Run(delay, func(t *testing.T) {
			f := SlowRequest(delay)
			start := time.Now()
			f(c)
			reqDuration := time.Since(start)
			expectedMinDuration, err := time.ParseDuration(delay)
			assert.NoError(t, err)
			assert.True(t, reqDuration > expectedMinDuration)
		})
	}
}
