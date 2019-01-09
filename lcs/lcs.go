package main

import (
	"fmt"
)

type Arr int

const (
	UP Arr = iota
	LEFT
	HIT
)

func Lcs(a, b string) ([][]int, [][]Arr) {
	n := len(a) + 1
	m := len(b) + 1
	ms := make([][]int, n)
	bs := make([][]Arr, n)
	for i := 0; i < n; i++ {
		ms[i] = make([]int, m)
		bs[i] = make([]Arr, m)
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if a[i-1] == b[j-1] {
				ms[i][j] = ms[i-1][j-1] + 1
				bs[i][j] = HIT
			} else if ms[i][j-1] < ms[i-1][j] {
				ms[i][j] = ms[i-1][j]
				bs[i][j] = UP
			} else {
				ms[i][j] = ms[i][j-1]
				bs[i][j] = LEFT
			}
		}
	}

	// var res int
	// for i := 0; i < n; i++ {
	// 	if res < ms[i][m-1] {
	// 		res = ms[i][m-1]
	// 	}
	// }

	// return nil, res
	return ms, bs
}

func printLCS(a string, ms [][]int, bs [][]Arr, i, j int) {
	if i == 0 || j == 0 {
		return
	} else if bs[i][j] == HIT {
		printLCS(a, ms, bs, i-1, j-1)
		fmt.Print(string(a[i-1]))
	} else if bs[i][j] == UP {
		printLCS(a, ms, bs, i-1, j)
	} else {
		printLCS(a, ms, bs, i, j-1)
	}
}

func main() {
	a := "ACCGGTCGAGTGCGCGGAAGCCGGCCGAA"
	b := "GTCGTTCGGAATGCCGTTGCTCTGTAAA"
	ms, bs := Lcs(a, b)
	fmt.Println(ms[len(a)][len(b)])
	printLCS(a, ms, bs, len(a), len(b))
}
