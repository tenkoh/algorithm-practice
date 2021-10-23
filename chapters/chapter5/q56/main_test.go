package q56

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		W int
		a []int
		m []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"want true", args{5, []int{1, 3}, []int{3, 2}}, true},
		{"want false", args{5, []int{1, 3}, []int{1, 2}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.W, tt.args.a, tt.args.m); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
