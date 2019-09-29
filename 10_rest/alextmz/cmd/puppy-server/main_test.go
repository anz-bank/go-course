package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	args = []string{"-d", "../../test/valid-formatted-json.json", "-p", "7735", "-s", "map"}
	// args = []string{""}
	os.Exit(m.Run())
}

func TestBadArguments(t *testing.T) {
	args = []string{"some", "invalid", "arguments"}
	// cmd := exec.Command(os.Args[0])
	assert.NotPanics(t, main)
	// os.Exit(0)
}

//     cmd.Env = append(os.Environ(), "TEST_MAIN=crasher")
//     err := cmd.Run()
//     if e, ok := err.(*exec.ExitError); ok && !e.Success() {
//         return
//     }
// 	t.Fatalf("process err %v, want exit status 1", err)

// func TestBadArgs(t *testing.T) {
// 	var err error
// 	//os.Exit(0)
// 	cmd := exec.Command(os.Args[0])
// 	out, err := cmd.CombinedOutput()
// 	sout := string(out) // because out is []byte
// 	if err != nil && !strings.Contains(sout, "somefunc failed") {
// 		fmt.Println(sout) // so we can see the full output
// 		t.Errorf("%v", err)
// 	}
// }

// func Test_main(t *testing.M) {
// 	var tests = map[string]struct {
// 		params   string
// 		expected string
// 	}{
// 		"Map store": {
// 			params: "-d ../.../test/valid-formatted-json.json -p 7735 -s map",
// 			expected: `Starting puppyserver with options:
// file  = ./test/valid-formatted-json.json
// port  = 7735
// store = map
// Loaded 3 puppies.
// `},
// 		"Sync store": {
// 			params: "-d ../.../test/valid-formatted-json.json -p 7735 -s sync",
// 			expected: `Starting puppyserver with options:
// file  = ./test/valid-formatted-json.json
// port  = 7735
// store = sync
// Loaded 3 puppies.
// `},
// 	}

// 	var buf bytes.Buffer
// 	out = &buf
// 	for k, v := range tests {
// 		args = []string{v.params}
// 		os.Exit(t.Run())
// 		actual := buf.String()
// 		assert.
// 		assert.Equal(t, tests[k].expected, actual)
// 	}
// }

// func Test_nofilegiven(t *testing.T) {
// 	var buf bytes.Buffer
// 	out = &buf
// 	args = []string{"-d", " "}
// 	main()
// 	actual := buf.String()
// 	assert.Equal(t, "open  : no such file or directory\n", actual)
// }

// func Test_invalidflag(t *testing.T) {
// 	var buf bytes.Buffer
// 	out = &buf
// 	args = []string{"-#"}
// 	main()
// 	actual := buf.String()
// 	assert.Equal(t, "unknown short flag '-#'", actual)
// }

// func Test_invalidpuppyID(t *testing.T) {
// 	var buf bytes.Buffer
// 	out = &buf
// 	args = []string{"-d", "../../test/invalid-ids.json"}
// 	main()
// 	actual := buf.String()
// 	assert.Equal(t, "400 Bad Request\n", actual)
// }

// func Test_invalidJSON(t *testing.T) {
// 	var buf bytes.Buffer
// 	out = &buf
// 	args = []string{"-d", "../../test/invalid-format.json"}
// 	main()
// 	actual := buf.String()
// 	assert.Equal(t, "json: cannot unmarshal object into Go value of type []puppy.Puppy\n", actual)
// }

// func Test_printpuppieserr(t *testing.T) {

// 	// puppies := []puppy.Puppy{{ID: 10}}
// 	puppystore := store.NewmapStore()
// 	printpuppies(puppystore, 1)

// 	var buf bytes.Buffer
// 	out = &buf
// 	args = []string{"-d", "../../test/invalid-format.json"}
// 	main()
// 	actual := buf.String()
// 	assert.Equal(t, "json: cannot unmarshal object into Go value of type []puppy.Puppy\n", actual)
// }

// func Test_readerror(t *testing.T) {
// 	var buf bytes.Buffer
// 	out = &buf
// 	args := "../../test/invalid-format.json"
// 	_, err := readfile(args, true)
// 	assert.Error(t, err)
// }
