package util

import (
	"testing"
)

func TestChmin(t *testing.T) {
	a := []int{2, 2, 2}
	type args struct {
		a *int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"want1", args{&a[0], 1}, 1},
		{"want2", args{&a[1], 3}, 2},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Chmin(tt.args.a, tt.args.b)
			if a[i] != tt.want {
				t.Errorf("want %d, got %d", tt.want, a[i])
			}
		})
	}
}
