package main

import (
	"bytes"
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	want := "[1 2 3 5]\n"

	t.Run("Test Main", func(t *testing.T) {
		var buf bytes.Buffer
		out = &buf
		main()
		result := buf.String()

		if result != want {
			t.Errorf("expected %#v, got %#v", want, result)
		}
	})
}

func Test_bubble(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want []int
	}{
		{"Bubblesort for {-3, -2, -1, -5}", []int{-3, -2, -1, -5}, []int{-5, -3, -2, -1}},
		{"Bubblesort for {-1}", []int{-1}, []int{-1}},
		{"Bubblesort for no parameters", []int{}, []int{}},
		{"Bubblesort for {0}", []int{0}, []int{0}},
		{"Bubblesort for {1}", []int{1}, []int{1}},
		{"Bubblesort for {0,1}", []int{0, 1}, []int{0, 1}},
		{"Bubblesort for {1,0}", []int{1, 0}, []int{0, 1}},
		{"Bubblesort for {3, 2, 1, 5}", []int{3, 2, 1, 5}, []int{1, 2, 3, 5}},
		{"Bubblesort for {1,1,2,2,9,4,1,0,1}", []int{1, 1, 2, 2, 9, 4, 1, 0, 1}, []int{0, 1, 1, 1, 1, 2, 2, 4, 9}},
	}

	for _, tr := range tests {
		tr2 := tr

		t.Run(tr2.name, func(t *testing.T) {

			origarg := make([]int, len(tr2.args))
			copy(origarg, tr2.args)

			result := bubble(tr2.args)

			if !reflect.DeepEqual(origarg, tr2.args) {
				t.Errorf("original slice was modified, expected %#v, got %#v", tr2.args, origarg)
			}
			if len(result) != len(tr2.want) {
				t.Errorf("wrong slice size, expected len=%d for %#v, got %d for %#v", len(tr2.want), tr2.want, len(result), result)
			}
			for k, v := range result {
				if v != tr2.want[k] {
					t.Errorf("wrong slice content, expected %#v, got %#v", tr2.want, result)
					break
				}
			}
		})
	}
}
