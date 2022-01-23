package main

import "fmt"

func partition(ints []int, l, r int) int {
	p := ints[r]
	for i := l; i < r; i++ {
		if ints[i] < p {
			ints[i], ints[l] = ints[l], ints[i]
			l++
		}
	}
	ints[l], ints[r] = ints[r], ints[l]
	return l
}

func quickSort(ints []int, l, r int) {
	if l > r {
		return
	}

	p := partition(ints, l, r)
	quickSort(ints, l, p-1)
	quickSort(ints, p+1, r)
}

func main() {
	ints := []int{8, 3, 1, 4, 5, 2}
	quickSort(ints, 0, len(ints)-1)
	fmt.Println(ints)
}

