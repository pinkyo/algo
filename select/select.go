package main

import (
	"../common"
	"../quicksort"
	"fmt"
	"math/rand"
	"time"
)


func Partition(A []int, p int, r int) int {
	x := A[r]
	i := p - 1
	for j := p; j <= r - 1; j++ {
		if A[j] <= x {
			i = i + 1
			A[i], A[j] = A[j], A[i]
		}
	}
	
	A[i + 1], A[r] = A[r], A[i + 1]
	return i + 1
}

func RandomizedPatition(A []int, p int, r int) int {
	i := rand.Intn(r - p + 1) + p
	A[i], A[r] = A[r], A[i]

	return Partition(A, p, r)
}

func RandomizedSelect(A []int, p int, r int, i int) int {
	if p == r {
		return A[p]
	}

	q := RandomizedPatition(A, p, r)
	k := q - p + 1
	if k == i {
		return A[q]
	} else if i < k {
		return RandomizedSelect(A, p, q - 1, i)
	} else {
		return RandomizedSelect(A, q + 1, r, i - k)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	data := common.MakeData(100000000)
	fmt.Println(RandomizedSelect(data, 0, len(data) - 1, 5000000))

	timeUsed := common.CalRunningTime(quicksort.RandomizedQuickSort, data)
	if common.VerifySortResult(data) {
		fmt.Println(timeUsed, " ns")
	} else {
		fmt.Println("Failed")
	}
	fmt.Println(data[5000000 - 1])
	// fmt.Println(RandomizedSelect(data, 0, len(data) - 1, 2))
	// fmt.Println(RandomizedSelect(data, 0, len(data) - 1, 3))
	// fmt.Println(RandomizedSelect(data, 0, len(data) - 1, 4))
	// fmt.Println(RandomizedSelect(data, 0, len(data) - 1, 5))
	// fmt.Println(RandomizedSelect(data, 0, len(data) - 1, 6))
	// fmt.Println(RandomizedSelect(data, 0, len(data) - 1, 7))
	// fmt.Println(RandomizedSelect(data, 0, len(data) - 1, 8))
	// fmt.Println(RandomizedSelect(data, 0, len(data) - 1, 9))
	// fmt.Println(RandomizedSelect(data, 0, len(data) - 1, 10))
}