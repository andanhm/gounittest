package tbt

import (
	"fmt"
	"testing"
)

func Test_Square(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		output  int
		wantErr bool
	}{
		{"square of 2 is 4", 2, 4, false},
		{"square of 3 is 9", 3, 9, false},
		{"square of 4 is 16", 4, 15, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			square := Square(tt.input)
			if (square != tt.output) != tt.wantErr {
				t.Errorf("TestSquare() output = %v, expected = %v", square, tt.output)
				t.Fail()
				return
			}
		})
	}
}
func TestSquare(t *testing.T) {
	tests := []struct {
		Input    int
		Expected int
	}{
		{2, 4},
		{3, 9},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Square(%d)", tt.Input), func(t *testing.T) {
			actual := Square(tt.Input)
			if actual != tt.Expected {
				t.Errorf("expected %d but got %d", tt.Expected, actual)
			}
		})
	}
}
