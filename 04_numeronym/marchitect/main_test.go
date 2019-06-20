package main

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)

func Test_numeronyms(t *testing.T) {
	tables := []struct {
		nameOfTestcase string
		input          []string
		output         []string
	}{
		{"normal test 1", []string{"what did you say", "cool;\nhah"}, []string{"w14y", "c7h"}},
		{"normal test 2", []string{"%^&*", "s**t", "?????"}, []string{"%2*", "s2t", "?3?"}},
		{"empty input", []string{}, []string{}},
		{"normal test 3", []string{"1234567", "1.23.", "666"}, []string{"157", "13.", "666"}},
		{"Mandarin Characters Hello", []string{"你好", "这个测试有意义", "谢谢你"}, []string{"你好", "这5义", "谢谢你"}},
	}
	for _, table := range tables {
		actual := numeronyms(table.input...)
		expected := table.output
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Unexpected output from test %s, expected = %s, actual = %s", table.nameOfTestcase, expected, actual)
		}
	}
}

func Test_main(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote(`[a11y K8s abc]
`)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main(), expected = %v, actual = %v", expected, actual)
	}
}
