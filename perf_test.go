package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const multiplier = 1000

var maxGoroutines = 8

func runBenchInsertSort(n int, b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	for N := 1; N <= maxGoroutines; N++ {
		b.Run(fmt.Sprintf("insert sort for %d threads", N), func(bn *testing.B) {
			for i := 0; i < bn.N; i++ {
				bn.StopTimer()
				arrayToSort := randomArray(n * multiplier)
				bn.StartTimer()
				sortInsert(arrayToSort, N)
			}
		})

	}
}

func runBenchHeapSort(n int, b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	for N := 1; N <= maxGoroutines; N++ {
		b.Run(fmt.Sprintf("heap sort for %d threads", N), func(bn *testing.B) {
			for i := 0; i < bn.N; i++ {
				bn.StopTimer()
				arrayToSort := randomArray(n * multiplier)
				bn.StartTimer()
				sortHeap(arrayToSort, N)
			}
		})

	}
}
func BenchmarkInsert25(b *testing.B)  { runBenchInsertSort(25, b) }
func BenchmarkInsert50(b *testing.B)  { runBenchInsertSort(50, b) }
func BenchmarkInsert100(b *testing.B) { runBenchInsertSort(100, b) }
func BenchmarkHeap25(b *testing.B)    { runBenchHeapSort(25, b) }
func BenchmarkHeap50(b *testing.B)    { runBenchHeapSort(50, b) }
func BenchmarkHeap100(b *testing.B)   { runBenchHeapSort(100, b) }
