package middleware

import (
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSlowRequestNegativeDuration(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	for i := -3; i >= 0; i++ {
		delay, err := time.ParseDuration(fmt.Sprintf("%ds", i))
		assert.NoError(t, err)
		expectedMaxDuration, err := time.ParseDuration("1s")
		t.Run(delay.String(), func(t *testing.T) {
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
		delay, err := time.ParseDuration(fmt.Sprintf("%ds", i))
		assert.NoError(t, err)
		t.Run(delay.String(), func(t *testing.T) {
			f := SlowRequest(delay)
			start := time.Now()
			f(c)
			reqDuration := time.Since(start)
			assert.True(t, reqDuration > delay)
		})
	}
}
