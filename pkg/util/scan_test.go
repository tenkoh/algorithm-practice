package util

import (
	"strings"
	"testing"
)

func TestScanInt(t *testing.T) {
	tests := []struct {
		input string
		want  []int
	}{
		{`1 2
		3 4`,
			[]int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		reader := strings.NewReader(tt.input)
		scan := NewScanner(reader)
		for _, want := range tt.want {
			if got := ScanInt(scan); got != want {
				t.Errorf("want %d, got %d\n", want, got)
			}
		}
	}
}
