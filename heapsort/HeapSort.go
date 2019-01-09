package main

import (
	"errors"
	"fmt"
	"../common"
)

type Heap struct {
	Data []int
	Length int
}

func Left(i int) int {
	return i << 1
}

func Right(i int) int {
	return i << 1 + 1
}

func Parent(i int) int {
	return i >> 1
}

func (heap *Heap) MaxHeapify(i int) {
	l := Left(i)
	r := Right(i)

	data := heap.Data
	var largest int

	if l <= heap.Length && data[l - 1] > data[i - 1] {
		largest = l
	} else {
		largest = i
	}

	if r <= heap.Length && data[r - 1] > data[largest - 1] {
		largest = r
	}

	if largest != i {
		data[largest - 1], data[i - 1] = data[i - 1], data[largest - 1]
		heap.MaxHeapify(largest)
	}
}

func (heap *Heap) BuildMaxHeap() {
	for i := heap.Length >> 1; i > 0; i-- {
		heap.MaxHeapify(i)
	}
}

func (heap *Heap) HeapMaximum() int {
	return heap.Data[0]
}

func (heap *Heap) HeapExtractMax() (int, error) {
	if heap.Length < 1 {
		return -1, errors.New("heap underflow")
	}
	
	data := heap.Data
	var max int
	max, data[0] = data[0], data[heap.Length - 1]
	heap.Length = heap.Length - 1

	heap.MaxHeapify(1)
	
	return max, nil
}

func (heap *Heap) HeapIncreaseKey(i int, key int) error {
	data := heap.Data
	if key < data[i] {
		return errors.New("new key is smaller than current key")
	}

	data[i - 1] = key
	for ;i > 1 && data[Parent(i) - 1] < data[i - 1]; {
		data[Parent(i) - 1], data[i - 1] = data[i - 1], data[Parent(i) - 1]
		i = Parent(i)
	}

	return nil
}

func (heap *Heap) MaxInsert(key int) error {
	data := heap.Data
	
	if len(data) == heap.Length { //full, allocate new slice
		newData := make([]int, len(data) * 2 + 1)
		copy(newData, data)
		data = newData
		heap.Data = newData
	}

	heap.Length += 1
	data[heap.Length - 1] = key - 1

	return heap.HeapIncreaseKey(heap.Length, key)
}

func (heap *Heap) PrintHeap() {
	fmt.Print("[")
	for i := 0; i < heap.Length; i++ {
		fmt.Print(heap.Data[i])
		if i != heap.Length - 1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("]")
}

func HeapSort(data []int) {
	heap := Heap{ data, len(data) }
	heap.BuildMaxHeap()
	for i := heap.Length; i >= 2; i-- {
		data[0], data[i - 1] = data[i - 1], data[0]
		heap.Length -= 1
		heap.MaxHeapify(1)
	}
}

func main() {
	data := common.MakeData(100000000)
	timeUsed := common.CalRunningTime(HeapSort, data)

	if common.VerifySortResult(data) {
		fmt.Println(timeUsed, " ns")
	} else {
		fmt.Println("Failed")
	}
}