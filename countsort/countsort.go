package main

import (
	"fmt"
	"../common"
)

func GetTenNum(data int, i int) int {
	temp := data
	for j := i; j > 0; j-- {
		temp /= 1000
	}

	return temp % 1000
}

func CountSort(A []int, k int) {
	C := make([]int, 1000)
	B := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		C[GetTenNum(A[i], k)] ++
	}

	for i := 1; i < len(C); i++ {
		C[i] += C[i -1]
	}

	for i := len(A) - 1; i >= 0; i-- {
		B[C[GetTenNum(A[i], k)] - 1] = A[i]
		C[GetTenNum(A[i], k)] --
	}

	for i := 0; i < len(A); i++ {
		A[i] = B[i]
	}
}

func RadixSort(A []int) { //A必须是正数, 且小于100000000
	max := 100000000;
	for i := 0; max > 0; i++ {
		CountSort(A, i)
		max /= 1000
	}
}

func main() {
	data := common.MakeData(100000000)
	timeUsed := common.CalRunningTime(RadixSort, data)
	if common.VerifySortResult(data) {
		fmt.Println(timeUsed, " ns")
	} else {
		fmt.Println("Failed")
	}
}