package intergers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expect := 4

	if sum != expect {
		t.Errorf("got %d want %d", sum, expect)
	}
}
