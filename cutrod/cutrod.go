package main

import "fmt"

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func extendedBottomUpCutRod(p []int, n int) (r []int, s []int) {
	r = make([]int, n+1)
	s = make([]int, n+1)
	r[0] = 0
	for j := 1; j <= n; j++ {
		q := INT_MIN
		for i := 1; i <= j; i++ {
			if q < p[i]+r[j-i] {
				q = p[i] + r[j-i]
				s[j] = i
			}
		}

		r[j] = q
	}

	return r, s
}

func printCutRodSolution(p []int, n int) {
	r, s := extendedBottomUpCutRod(p, n)

	fmt.Println(r[n])
	for n > 0 {
		fmt.Print(s[n], " ")
		n = n - s[n]
	}
}

func main() {
	p := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	printCutRodSolution(p, 6)
}
