package main

import (
	"fmt"
	"maps"
	"slices"
)

func Filter[T any](s []T, keep func(T) bool) []T {
	result := make([]T, 0, len(s))
	for _, v := range s {
		if keep(v) {
			result = append(result, v)
		}
	}
	return result
}

func Keys[K comparable, V any](m map[K]V) []K {
	return slices.Collect(maps.Keys(m))
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	even := Filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println(even)

	m := map[string]int{"go": 1, "rust": 2, "zig": 3}
	keys := Keys(m)
	slices.Sort(keys)
	fmt.Println(keys)
}
