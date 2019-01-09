package common

import (
	"fmt"
	"math/rand"
	"time"
)

func VerifySortResult(result []int) bool {
	for i := 1; i < len(result); i++ {
		if result[i] < result[i - 1] {
			return false
		}
	}

	return true;
}

func MakeData(length int) []int {
	rand.Seed(time.Now().UnixNano())
	
	data := make([]int, length)
	for i := 0; i < length; i++ {
		data[i] = rand.Int()
	}
	return data;
}

func CalRunningTime(Sort func([]int), data []int) int64 {
	start := time.Now().UnixNano()
	Sort(data)
	return (time.Now().UnixNano() - start)
}

func RunAndVerifySortFunction(Sort func([]int), maxRunCount int) {
	successCount := 0
	for c := 0; c < maxRunCount; c++ {
		data := MakeData(100)
		Sort(data)
		if VerifySortResult(data) {
			successCount ++
		}
	}

	fmt.Printf("Total Run %d, Success %d.\n", maxRunCount, successCount)
}