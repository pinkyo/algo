package main

import "fmt"

/*
01 knapsack problem, using dynamic programing.
*/
func knapsack(w, v []int, W int) ([]int, []int) {
	max := make([]int, W+1)
	pre := make([]int, W+1)

	for i := 0; i <= W; i++ {
		pre[i] = -1
		max[i] = 0
	}

	size := len(w)
	for i := 0; i < size; i++ {
		weight := w[i]
		value := v[i]

		for j := W; j >= weight; j-- {
			if max[j] < max[j-weight]+value {
				max[j] = max[j-weight] + value
				pre[j] = i
			}
		}
	}

	return max, pre
}

func printSolution(w, v, max, pre []int) {
	W := len(max) - 1
	fmt.Println("Max value is ", max[W])

	fmt.Println("Solution:")
	item := pre[W]
	left := W
	for item != -1 {
		fmt.Println("(", w[item], ",", v[item], ")")
		left = left - w[item]
		item = pre[left]
	}
}

func main() {
	w := []int{10, 20, 30}
	v := []int{60, 100, 120}

	max, pre := knapsack(w, v, 50)
	printSolution(w, v, max, pre)
}
