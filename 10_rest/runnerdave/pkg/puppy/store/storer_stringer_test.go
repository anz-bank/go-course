package store

import (
	"strings"
	"testing"

	tassert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStringerMethod(t *testing.T) {
	// given
	s := NewSyncStore()
	assert := tassert.New(t)
	testPuppy1 := puppy1()
	testPuppy2 := puppy2()
	createError1 := s.CreatePuppy(testPuppy1)
	createError2 := s.CreatePuppy(testPuppy2)
	r := require.New(t)
	r.NoError(createError1, "Create should not produce an error")
	r.NoError(createError2, "Create should not produce an error")

	// when
	expected := `{"breed":"Chihuahua","color":"Brown","id":1,"value":12.3} 
	{"breed":"Cacri","color":"Undefined","id":2,"value":1.3} `
	actual := s.String()

	// then
	assert.Equal(strings.Replace(expected, "\n\t", "", -1), actual)
}
