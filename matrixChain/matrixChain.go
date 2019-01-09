package main

import "fmt"

const INT_MAX = int(^uint(0) >> 1)

func MatrixChainOrder(p []int) int {
	n := len(p) - 1
	m := make([][]int, n)
	s := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
		s[i] = make([]int, n)
		m[i][i] = 0
	}

	for l := 1; l < n; l++ {
		for i := 0; i < n-l; i++ {
			j := i + l
			m[i][j] = INT_MAX
			for k := i; k <= j-1; k++ {
				q := m[i][k] + m[k+1][j] + p[i]*p[k+1]*p[j+1]
				if q < m[i][j] {
					m[i][j] = q
					s[i][j] = k
				}
			}
		}
	}

	return m[0][n-1]
}

func main() {
	// p := []int{10, 100, 5, 50}
	p := []int{30, 35, 15, 5, 10, 20, 25}
	res := MatrixChainOrder(p)
	fmt.Println(res)
}
