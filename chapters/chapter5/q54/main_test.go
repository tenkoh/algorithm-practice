package q54

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		W int
		k int
		a []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"want true", args{5, 2, []int{1, 2, 3}}, true},
		{"want false", args{6, 2, []int{1, 2, 3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.W, tt.args.k, tt.args.a); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
