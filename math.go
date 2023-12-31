package genericgokit

import (
	"golang.org/x/exp/constraints"
)

func Min[T constraints.Ordered](collection []T) T {
	var min T
	if len(collection) == 0 {
		return min
	}
	min = collection[0]
	for i := 1; i < len(collection); i++ {
		item := collection[i]

		if item < min {
			min = item
		}
	}
	return min
}

func Max[T constraints.Ordered](collection []T) T {
	var max T
	if len(collection) == 0 {
		return max
	}
	max = collection[0]
	for i := 1; i < len(collection); i++ {
		item := collection[i]
		if item > max {
			max = item
		}
	}
	return max
}
