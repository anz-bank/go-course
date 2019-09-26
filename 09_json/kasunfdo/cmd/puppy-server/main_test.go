package main

import (
	"bytes"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainOutputNoErrorLongFlag(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	os.Args = []string{"", "--data", "../../data/data.json"}
	main()

	expected := `{0 Labrador White 1200} (id: 1) added to store
{0 Boxer Black 2099.99} (id: 2) added to store
{0 Terrier Brown 1099.99} (id: 3) added to store
{0 Husky Black & White 1999.99} (id: 4) added to store
{0 Retriever Gold 3099.99} (id: 5) added to store
`
	actual := buf.String()
	assert.Equal(t, expected, actual)
}

func TestMainOutputNoErrorShotFlag(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	os.Args = []string{"", "-d", "../../data/data.json"}
	main()

	expected := `{0 Labrador White 1200} (id: 1) added to store
{0 Boxer Black 2099.99} (id: 2) added to store
{0 Terrier Brown 1099.99} (id: 3) added to store
{0 Husky Black & White 1999.99} (id: 4) added to store
{0 Retriever Gold 3099.99} (id: 5) added to store
`
	actual := buf.String()
	assert.Equal(t, expected, actual)
}

func TestMainInvalidFlag(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	os.Args = []string{"", "--foo", "../../data/data.json"}
	assert.Panics(t, main)
}

func TestMainNonExistingFile(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	os.Args = []string{"", "--data", "../../data/foo.json"}
	assert.Panics(t, main)
}

func TestMainInvalidDataType(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	os.Args = []string{"", "--data", "../../data/invalidData1.json"}
	assert.Panics(t, main)
}

func TestMainInvalidPuppyData(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	os.Args = []string{"", "--data", "../../data/invalidData2.json"}
	assert.Panics(t, main)
}

type mockInvalidReader string

func (mockInvalidReader) Read(_ []byte) (n int, err error) {
	return 0, errors.New("mockInvalidReader:error")
}

func TestCreatePuppiesReadFail(t *testing.T) {
	dataFile := mockInvalidReader("")
	err := createPuppies(dataFile, nil)
	assert.Error(t, err)
	assert.Equal(t, "internal error", err.Error())
}
