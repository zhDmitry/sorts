package main

import "sort"

func insertionsort(array []int) []int {
	var n = len(array)
	for i := 1; i < n; i++ {
		x := array[i]
		j := i - 1

		loc := sort.SearchInts(array[:i], x)
		for j >= loc {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = x
	}
	return array
}
