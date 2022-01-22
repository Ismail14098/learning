package main

import "fmt"

func binarySearchRecursive(ints []int, l, r, target int) int {
	if l == r {
		return -1
	}

	i := (l+r)/2
	if ints[i] == target {
		return i
	} else if target > ints[i] {
		return binarySearchRecursive(ints, i+1, r, target)
	} else {
		return binarySearchRecursive(ints, l, i-1, target)
	}
}

func binarySearch(ints []int, target int) int {
	l := 0
	r := len(ints) - 1
	for l <= r {
		if l == r {
			return -1
		}
		mid := (l+r)/2
		if ints[mid] == target {
			return mid
		} else if target > ints[mid] {
			l = mid+1
		} else {
			r = mid-1
		}
	}
	return -1
}

func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(binarySearchRecursive(ints, 0, len(ints)-1, 4))
	fmt.Println(binarySearch(ints, 4))
}
