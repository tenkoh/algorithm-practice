package q51

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		N int
		a []int
		b []int
		c []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test", args{5, []int{5, 3, 1, 5, 5}, []int{3, 5, 3, 3, 3}, []int{1, 1, 5, 1, 1}}, 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.N, tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
