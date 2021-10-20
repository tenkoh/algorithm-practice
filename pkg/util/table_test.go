package util

import (
	"reflect"
	"testing"
)

func TestMakeTable(t *testing.T) {
	type args struct {
		row  int
		col  int
		init int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"test1", args{2, 3, 0}, [][]int{{0, 0, 0}, {0, 0, 0}}},
		{"test1", args{2, 3, 1}, [][]int{{1, 1, 1}, {1, 1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeTable(tt.args.row, tt.args.col, tt.args.init); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
