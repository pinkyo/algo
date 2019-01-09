package main

import (
	"fmt"
	"./common"
)

func InsertionSort(data []int) {
	for j := 1; j < len(data); j++ {
		key := data[j]
		i := j - 1
		for i >=  0 && data[i] > key {
			data[i + 1] = data[i]
			i = i - 1
		}
		data[i + 1] = key
	}
}

func MergeSort(data []int) {
	MergeSortHelper(data, 0, len(data) - 1)
}

func MergeSortHelper(data []int, p int, r int) {
	if (p < r) {
		q := ( p + r ) / 2
		MergeSortHelper(data, p, q)
		MergeSortHelper(data, q + 1, r)
		Merge(data, p, q, r)
	}
}

func Merge(data []int, p int, q int, r int) {
	x := make([]int, q - p + 1)
	y := make([]int, r - q)

	copy(x, data[p:])
	copy(y, data[q+1:])

	i, j := 0, 0
	for k := p; k <= r; k++ {
		if (i == len(x)) {
			data[k] = y[j]
			j++
		} else if (j == len(y)) {
			data[k] = x[i]
			i++
		} else if x[i] <= y[j] {
			data[k] = x[i]
			i++
		} else {
			data[k] = y[j]
			j++
		}
	}
}

func Reverse(data []int) {
	length := len(data)
	for i := 0; i < length / 2; i++ {
		data[i], data[length - i - 1] = data[length - i - 1], data[i]
	}
}

func main() {
	data := common.MakeData(100000000)
	timeUsed := common.CalRunningTime(MergeSort, data)
	if common.VerifySortResult(data) {
		fmt.Println(timeUsed, " ns")
	} else {
		fmt.Println("Failed")
	}
}
