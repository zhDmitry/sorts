package main

import (
	"reflect"
	"testing"
)

func TestInsertSort(t *testing.T) {
	arrayToSort := []int{5, 6, 4, 2, 3, 1}
	res := sortInsert(arrayToSort, 1)
	expectedRes := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(res, expectedRes) {
		t.Fatalf("array is not sorted, expected %v but got %v", expectedRes, res)
	}
}
func TestHeapSort(t *testing.T) {
	arrayToSort := []int{5, 6, 4, 2, 3, 1}
	res := sortHeap(arrayToSort, 1)
	expectedRes := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(res, expectedRes) {
		t.Fatalf("array is not sorted, expected %v but got %v", expectedRes, res)
	}
}
