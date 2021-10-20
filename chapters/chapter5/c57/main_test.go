package c57

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		N      int
		W      int
		weight []int
		value  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{6, 15, []int{2, 1, 3, 2, 1, 5}, []int{3, 2, 6, 1, 3, 85}}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.N, tt.args.W, tt.args.weight, tt.args.value); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
