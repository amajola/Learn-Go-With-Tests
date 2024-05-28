package iteration

import "testing"

func TestRepeat(t *testing.T) {
	reapeated := Repeat("a")
	expected := "aaaaa"

	if reapeated != expected {
		t.Errorf("expected %q got %q", reapeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
