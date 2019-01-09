package quicksort

import (
	"fmt"
	"math/rand"
	"../common"
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

func HoarePartition(A []int, p int, r int) int {
	x := A[p]

	i, j := p - 1, r + 1
	for true {
		for {
			i++
			if i == r || A[i] >= x {
				break
			}
		}
		
		for {
			j--
			if j == p || A[j] <= x {
				break
			}
		}

		if i < j {
			A[i], A[j] = A[j], A[i]
		} else {
			return j
		}
	}

	return j
}

func RandomizedPartition(A []int, p int, r int) int {
	i :=  rand.Intn(r - p + 1) + p
	A[i], A[r] = A[r], A[i]
	return HoarePartition(A, p, r)
}

func QuickSortHelper(A []int, p int, r int) {
	if p < r {
		q := Partition(A, p, r)
		QuickSortHelper(A, p, q - 1)
		QuickSortHelper(A, q + 1, r)
	}
}

func RandomizedQuickSortHelper(A []int, p int, r int) {
	for p < r {
		q := RandomizedPartition(A, p, r)
		RandomizedQuickSortHelper(A, p, q)
		// RandomizedQuickSortHelper(A, q + 1, r)
		p = q + 1
	}
}

func QuickSort(A []int) {
	QuickSortHelper(A, 0, len(A) - 1)
}

func RandomizedQuickSort(A []int) {
	RandomizedQuickSortHelper(A, 0, len(A) - 1)
}

func main() {
	// common.RunAndVerifySortFunction(RandomizedQuickSort, 100)

	// data := common.MakeData(10)
	// fmt.Println(data)
	// RandomizedQuickSort(data)
	// fmt.Println(data)

	data := common.MakeData(100000000)
	timeUsed := common.CalRunningTime(RandomizedQuickSort, data)
	if common.VerifySortResult(data) {
		fmt.Println(timeUsed, " ns")
	} else {
		fmt.Println("Failed")
	}
}