package q55

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		W int
		a []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"want true", args{10, []int{2, 3, 4}}, true},
		{"want false", args{2, []int{3, 4}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.W, tt.args.a); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
