package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote(`Puppy created in map of Breed:Chihuahua, errors at creation:<nil>
Puppy created in sync of Breed:Chihuahua, value updated to:10.300000, error at creation:<nil>, error in update:<nil>
2: puppy with ID:12 not found`)
	actual := strconv.Quote(buf.String())
	t.Logf("expected:%s", expected)
	t.Logf("actual:%s", actual)

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}
