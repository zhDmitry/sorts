package main

import (
	"flag"
	"math/rand"
	"time"
)

func Merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}

func randomArray(len int) []int {
	a := make([]int, len)
	for i := 0; i <= len-1; i++ {
		a[i] = rand.Intn(len)
	}
	return a
}

func runThreads(array []int, results chan<- []int, N int) {
	size := len(array)
	itemsPerThread := size / N
	for i := 0; i < N; i++ {
		start := i * itemsPerThread
		var end int
		if i == N-1 {
			end = len(array)
		} else {
			end = (i + 1) * itemsPerThread

		}
		go func(start, end int) {
			results <- insertionsort(array[start:end])
		}(start, end)
	}
}

func runThreadsHeapSort(array []int, results chan<- []int, N int) {
	var heap = new(Heap)
	size := len(array)
	itemsPerThread := size / N
	for i := 0; i < N; i++ {
		start := i * itemsPerThread
		var end int
		if i == N-1 {
			end = len(array)
		} else {
			end = (i + 1) * itemsPerThread

		}
		go func(start, end int) {
			arr := array[start:end]
			heap.HeapSort(arr)
			results <- arr
		}(start, end)
	}
}

func runMergeBinary(results chan []int, initialArrLen int) []int {
	var result []int
	for arr1 := range results {
		if len(arr1) == initialArrLen {
			result = arr1
			break
		}
		arr2 := <-results
		go func(arr1, arr2 []int) {
			results <- Merge(arr1, arr2)
		}(arr1, arr2)

	}
	return result

}

func sortInsert(array []int, N int) []int {
	results := make(chan []int, 2)
	initialArrLen := len(array)
	runThreads(array, results, N)
	return runMergeBinary(results, initialArrLen)
}

func sortHeap(array []int, N int) []int {
	results := make(chan []int, 2)
	initialArrLen := len(array)
	runThreadsHeapSort(array, results, N)
	return runMergeBinary(results, initialArrLen)
}

var n = *flag.Int("n", 100, "number size")
var nt = *flag.Int("N", 5, "number of threads")

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	arrayToSort := randomArray(n)
	sortHeap(arrayToSort, nt)
}
