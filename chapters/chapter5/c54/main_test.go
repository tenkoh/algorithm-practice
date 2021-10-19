package c54

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		N int
		h []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{7, []int{2, 9, 4, 5, 1, 6, 10}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.N, tt.args.h); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
