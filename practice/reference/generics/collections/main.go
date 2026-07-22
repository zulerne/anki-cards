package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	values := []int{4, 1, 3, 2}
	slices.Sort(values)

	original := map[string]int{"go": 1}
	clone := maps.Clone(original)
	clone["cards"] = 2

	fmt.Println(values)
	fmt.Println(original, clone)
}
