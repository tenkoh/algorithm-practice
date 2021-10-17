package q45

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		N     int
		count int
	}{
		{1000, 6},
		{10000, 3*10*6 + 9*6},
	}

	for _, tt := range tests {
		if got := Solution(tt.N); got != tt.count {
			t.Errorf("want %d, got %d\n", tt.count, got)
		}
	}
}
