package q45

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		N     int
		count int
	}{
		{1000, 6},
		{10000, 6 + 3*6*2},
	}

	for _, tt := range tests {
		if got := Solution(tt.N); got != tt.count {
			t.Errorf("want %d, got %d\n", tt.count, got)
		}
	}
}

func TestBinary(t *testing.T) {
	if 2 != 0b10 {
		t.Error("not equal")
	}
	if 0b10|0b01 != 3 {
		t.Error("not equal")
	}
	if 1<<1 != 2 {
		t.Error("not equal")
	}
	if 1<<0 != 1 {
		t.Error("not equal")
	}
}
