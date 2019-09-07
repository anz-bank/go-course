package main

import (
	"bytes"
	"strconv"
	"strings"
	"testing"
)

func TestMainOutputMap(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	args = []string{"--data", "../../puppydata/data.json", "--store", "map", "--test", "true", "-p", "3005"}

	go main()
	<-c

	readable := strconv.Quote(`Map store of puppies:
	{map[1:{ULTRASURE green 1 2732.81} 
	2:{EPLODE brown 2 3889.92} 
	3:{SENTIA blue 3 1472.29} 
	4:{MAKINGWAY brown 4 2254.7} 
	5:{VISUALIX brown 5 3250.57} 
	6:{SPLINX brown 6 2467.04} 
	7:{POWERNET blue 7 3970.65}] 8}`)
	expected := strings.ReplaceAll(readable, "\\n\\t", "")
	actual := strconv.Quote(buf.String())
	t.Logf("expected:%s", expected)
	t.Logf("actual:%s", actual)

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestMainOutputSync(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	args = []string{"--data", "../../puppydata/data.json", "--store", "sync", "--test", "true", "-p", "3001"}

	go main()
	<-c

	readable := strconv.Quote(`Sync store of puppies:
	{"breed":"ULTRASURE","color":"green","id":1,"value":2732.81} 
	{"breed":"EPLODE","color":"brown","id":2,"value":3889.92} 
	{"breed":"SENTIA","color":"blue","id":3,"value":1472.29} 
	{"breed":"MAKINGWAY","color":"brown","id":4,"value":2254.7} 
	{"breed":"VISUALIX","color":"brown","id":5,"value":3250.57} 
	{"breed":"SPLINX","color":"brown","id":6,"value":2467.04} 
	{"breed":"POWERNET","color":"blue","id":7,"value":3970.65} `)
	expected := strings.ReplaceAll(readable, "\\n\\t", "")
	actual := strconv.Quote(buf.String())
	t.Logf("expected:%s", expected)
	t.Logf("actual:%s", actual)

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}

}

func TestMainOutputInvalidStorageFlag(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	args = []string{"--data", "../../puppydata/data.json", "-s", "mysql", "--test", "true", "-p", "3000"}
	go main()
	<-c

	actual := buf.String()
	expected := "Invalid storage"
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

func TestMainOutputInvalidValuePuppy(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	args = []string{"--data", "../../puppydata/invalid-puppies-negative-value.json", "-t", "true"}

	go main()
	<-c

	readable := strconv.Quote(`Could not create puppy, error: 1: 
	puppy has invalid value (-2732.810059)
	Could not setup sync storage:1: puppy has invalid value (-2732.810059)
	Sync store of puppies:`)
	expected := strings.ReplaceAll(readable, "\\n\\t", "")
	actual := strconv.Quote(buf.String())
	t.Logf("expected:%s", expected)
	t.Logf("actual:%s", actual)

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestMainOutputInvalidValuePuppyMapStorage(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	args = []string{"--data", "../../puppydata/invalid-puppies-negative-value.json", "-s", "map", "-t", "true",
		"-p", "3050"}

	go main()
	<-c

	readable := strconv.Quote(`Could not create puppy, error: 1: 
	puppy has invalid value (-2732.810059)
	Could not setup map storage:1: puppy has invalid value (-2732.810059)Map store of puppies:{map[] 0}`)
	expected := strings.ReplaceAll(readable, "\\n\\t", "")
	actual := strconv.Quote(buf.String())
	t.Logf("expected:%s", expected)
	t.Logf("actual:%s", actual)

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestInvalidPort(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	args = []string{"--data", "../../puppydata/data.json", "-s", "map",
		"-p", "w"}

	main()

	readable := strconv.Quote(`Map store of puppies:
	{map[1:{ULTRASURE green 1 2732.81} 
	2:{EPLODE brown 2 3889.92} 
	3:{SENTIA blue 3 1472.29} 
	4:{MAKINGWAY brown 4 2254.7} 
	5:{VISUALIX brown 5 3250.57} 
	6:{SPLINX brown 6 2467.04} 
	7:{POWERNET blue 7 3970.65}] 8}
	Could not start puppy server: listen tcp: 
	lookup tcp/w: nodename nor servname provided, or not known`)
	expected := strings.ReplaceAll(readable, "\\n\\t", "")
	actual := strconv.Quote(buf.String())
	t.Logf("expected:%s", expected)
	t.Logf("actual:%s", actual)

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}
