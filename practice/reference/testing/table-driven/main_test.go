package tabledriven

import "testing"

func add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{name: "positive", a: 1, b: 2, want: 3},
		{name: "negative", a: -1, b: -2, want: -3},
		{name: "zero", want: 0},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if got := add(test.a, test.b); got != test.want {
				t.Fatalf("add(%d, %d) = %d, want %d", test.a, test.b, got, test.want)
			}
		})
	}
}
