package main

import (
	"testing"
)

func add(a, b int) int {
	return a + b
}
func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{name: "positive", a: 2, b: 3, want: 5},
		{name: "negative", a: -1, b: -3, want: -4},
		{name: "zero", a: 0, b: 0, want: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := add(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
func main() {
}
