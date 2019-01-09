package main

import "fmt"

func gas(s, f []int) []bool {
	n := len(s)
	result := make([]bool, n)

	result[0] = true
	k := 0
	for i := 1; i < n; i++ {
		if s[i] >= f[k] {
			result[i] = true
			k = i
		}
	}

	return result
}

func printlnGasResult(result []bool) {
	for i := 0; i < len(result); i++ {
		if result[i] {
			fmt.Print("A", i+1, ", ")
		}
	}
	fmt.Println()
}

func main() {
	s := []int{1, 3, 0, 5, 3, 5, 6, 8, 8, 2, 12}
	f := []int{4, 5, 6, 7, 9, 9, 10, 11, 12, 14, 16}
	// fmt.Println(gas(s, f))
	result := gas(s, f)
	printlnGasResult(result)
}
