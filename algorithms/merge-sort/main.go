package main

import (
	"fmt"
)

func merge(left, right []int) []int {
	result := make([]int, len(left) + len(right))
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}
	for i := i; i < len(left); i++ {
		result[k] = left[i]
		k++
	}
	for j := j; j < len(right); j++ {
		result[k] = right[j]
		k++
	}
	return result
}

func mergeSort(ints []int, l, r int) []int {
	if l == r {
		return []int{ints[l]}
	}

	m := (l+r)/2
	left := mergeSort(ints, l, m)
	right := mergeSort(ints, m+1, r)
	return merge(left, right)
}

func main() {
	ints := []int{8, 3, 1, 4, 5, 2}
	fmt.Println(mergeSort(ints, 0, len(ints)-1))
}
