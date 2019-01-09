package main

import (
	"fmt"
	"time"
)

func fibonacci(n int) int {
	r := make([]int, n+1)
	r[0] = 1
	r[1] = 1
	for i := 2; i <= n; i++ {
		r[i] = r[i-1] + r[i-2]
	}
	return r[n]
}

func fibonacciR(n int) int {
	r := make([]int, n+1)
	return fibonacciRecur(n, r)
}

func fibonacciRecur(n int, r []int) int {
	if n <= 1 {
		return 1
	} else if r[n] == 0 {
		r[n] = fibonacciRecur(n-1, r) + fibonacciRecur(n-2, r)
		if r[n] < 0 {
			fmt.Println("Overflow")
		}
	}

	return r[n]
}

func main() {
	start := time.Now().UnixNano()
	fmt.Println(fibonacciR(70))
	timeUsed := (time.Now().UnixNano() - start)

	fmt.Println("Time Used: ", timeUsed)
}
