package main

import "testing"

func add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{
			name: "add 1",
			a:    1,
			b:    2,
			want: 3,
		},
		{
			name: "add 2",
			a:    -1,
			b:    -2,
			want: -3,
		},
		{
			name: "add 3",
			a:    0,
			b:    0,
			want: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := add(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}

}
func main() {}
