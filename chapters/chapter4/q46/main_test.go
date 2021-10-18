package q46

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		a []int
		W int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{3, 4, 5, 6}, 4}, true},
		{"test1", args{[]int{3, 4, 5, 6}, 10}, true},
		{"test1", args{[]int{3, 4, 5, 6}, 19}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.a, tt.args.W); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
