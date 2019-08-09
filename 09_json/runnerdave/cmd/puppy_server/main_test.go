package main

import (
	"bytes"
	"strconv"
	"strings"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	args = []string{"--data", "../../puppydata/data.json"}

	main()

	readable := strconv.Quote(`Map store of puppies:
	{map[2:{ULTRASURE green 2 2732.81} 
	3:{EPLODE brown 3 3889.92} 
	4:{SENTIA blue 4 1472.29} 
	5:{MAKINGWAY brown 5 2254.7} 
	6:{VISUALIX brown 6 3250.57} 
	7:{SPLINX brown 7 2467.04} 
	8:{POWERNET blue 8 3970.65}]}
	
	Sync store of puppies:
	{"breed":"ULTRASURE","color":"green","id":2,"value":2732.81} 
	{"breed":"EPLODE","color":"brown","id":3,"value":3889.92} 
	{"breed":"SENTIA","color":"blue","id":4,"value":1472.29} 
	{"breed":"MAKINGWAY","color":"brown","id":5,"value":2254.7} 
	{"breed":"VISUALIX","color":"brown","id":6,"value":3250.57} 
	{"breed":"SPLINX","color":"brown","id":7,"value":2467.04} 
	{"breed":"POWERNET","color":"blue","id":8,"value":3970.65} `)
	expected := strings.ReplaceAll(readable, "\\n\\t", "")
	actual := strconv.Quote(buf.String())
	t.Logf("expected:%s", expected)
	t.Logf("actual:%s", actual)

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestMainOutputNoValueForFlag(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	args = []string{"-d", " "}
	main()

	actual := buf.String()
	expected := "Command line could not be parsed, error: path ' ' does not exist"
	t.Logf("expected:%s", expected)
	t.Logf("actual:%s", actual)
	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestMainOutputInvalidData(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	args = []string{"--data", "../../puppydata/invalid-data.json"}

	main()

	expected := strconv.Quote("Could not unmarshall puppies, error: invalid character 'r' looking for beginning of value")
	actual := strconv.Quote(buf.String())
	t.Logf("expected:%s", expected)
	t.Logf("actual:%s", actual)

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestMainOutputInvalidPuppy(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	args = []string{"--data", "../../puppydata/invalid-puppies.json"}

	main()

	readable := strconv.Quote(`Could not unmarshall puppies, 
	error: json: cannot unmarshal string into Go struct field Puppy.id of type int16`)
	expected := strings.ReplaceAll(readable, "\\n\\t", "")
	actual := strconv.Quote(buf.String())
	t.Logf("expected:%s", expected)
	t.Logf("actual:%s", actual)

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestMainOutputRepeatedPuppy(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	args = []string{"--data", "../../puppydata/repeated-puppy.json"}

	main()

	readable := strconv.Quote(`Could not create puppy, error: 0: puppy with id 2 already exists
	Map store of puppies:{map[]}
	Could not create puppy, error: 0: puppy with id 2 already exists
	Sync store of puppies:{"breed":"ULTRASURE","color":"green","id":2,"value":2732.81} `)
	expected := strings.ReplaceAll(readable, "\\n\\t", "")
	actual := strconv.Quote(buf.String())
	t.Logf("expected:%s", expected)
	t.Logf("actual:%s", actual)

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}
