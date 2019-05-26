package fibo

import (
	"testing"
)

func TestFibo(t *testing.T) {
	tables := []struct {
		x int
		n int
	}{
		{1, 1},
		{7, 13},
		{2, 1},
		{3, 2},
		{6, 8},
	}

	for _, table := range tables {
		total := fib(table.x)
		if total != table.n {
			t.Errorf("Fibo of (%d) was incorrect, got: %d, want: %d.", table.x, total, table.n)
		}
	}
}

func TestFibo7(t *testing.T) {
	val := fib(7)
	if val != 13 {
		t.Errorf("Should have returned 13, but returned: (%d) instead", val)
	}
}
