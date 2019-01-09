package main

import (
	"fmt"
	"time"
	"../common"
)

func FindMaxCrossingSubarray(data []int, low, mid, high int) (int, int, int) {
	leftSum := data[mid]
	maxLeft := mid
	sum := data[mid]
	for i := mid - 1; i >= low; i-- {
		sum += data[i]
		if leftSum < sum {
			leftSum = sum
			maxLeft = i
		}
	}

	rightSum := data[mid + 1]
	maxRight := mid + 1
	sum = data[mid + 1]
	for i := mid + 2; i <= high; i++ {
		sum += data[i]
		if rightSum < sum {
			rightSum = sum
			maxRight = i
		}
	}

	return maxLeft, maxRight, leftSum + rightSum
}

func FindMaxSubarray(data []int, low, high int) (int, int, int) {
	if low == high {
		return low, high, data[low]
	}
	mid := ( low + high ) / 2
	leftLow, leftHigh, leftSum := FindMaxSubarray(data, low, mid)
	rightLow, rightHigh, rightSum := FindMaxSubarray(data, mid + 1, high)
	crossLow, crossHigh, crossSum := FindMaxCrossingSubarray(data, low, mid, high)

	if leftSum >= rightSum && leftSum >= crossSum {
		return leftLow, leftHigh, leftSum
	} else if rightSum >= leftSum && rightSum >= crossSum {
		return rightLow, rightHigh, rightSum
	} else {
		return crossLow, crossHigh, crossSum
	}
}

func main() {
	data := common.MakeData(100000000);
	start := time.Now().UnixNano()
	low, high, sum := FindMaxSubarray(data, 0, len(data) - 1)
	fmt.Println(time.Now().UnixNano() - start, "ns")
	fmt.Println(low, high, sum)
}