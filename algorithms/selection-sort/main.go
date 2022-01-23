package main

import "fmt"

func findLowest(ints []int) int {
	lowest := ints[0]
	index := 0
	for i := 0; i < len(ints); i++ {
		if ints[i] < lowest {
			lowest = ints[i]
			index = i
		}
	}
	return index
}

func deleteElement(ints[]int, l int) []int {
	Len := len(ints)-1
	result := make([]int, Len, Len)
	k := 0
	for i := 0; i < len(ints); i++ {
		if i != l {
			result[k] = ints[i]
			k++
		}
	}
	return result
}

func selectionSort(ints []int) []int {
	var result []int
	max := len(ints)
	for i := 0; i < max; i++ {
		l := findLowest(ints)
		result = append(result, ints[l])
		newints := deleteElement(ints, l)
		ints = newints
	}
	return result
}

func main() {
	ints := []int{8, 3, 1, 4, 5, 2}
	fmt.Println(selectionSort(ints))
}
