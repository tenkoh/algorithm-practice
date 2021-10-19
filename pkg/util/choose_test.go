package util

import (
	"testing"
)

func TestChooseInt(t *testing.T) {
	a := []int{2, 2, 2}
	type args struct {
		a    *int
		b    int
		mode string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"want1", args{&a[0], 1, "min"}, 1},
		{"want2", args{&a[1], 3, "min"}, 2},
		{"want3", args{&a[2], 3, "max"}, 3},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ChooseInt(tt.args.a, tt.args.b, tt.args.mode)
			if a[i] != tt.want {
				t.Errorf("want %d, got %d", tt.want, a[i])
			}
		})
	}
}
