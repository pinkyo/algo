package main

import (
	"errors"
	"fmt"
)

func MultiplySquareMatrix(a, b [][]int) ([][]int, error) {
	if len(a[0]) != len(b) {
		return nil, errors.New("Incompatible dimensions")
	}

	n := len(a)
	m := len(b[0])
	aw := len(b)
	c := make([][]int, n)

	for i := 0; i < n; i++ {
		c[i] = make([]int, m)
		for j := 0; j < m; j++ {
			c[i][j] = 0
			for k := 0; k < aw; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return c, nil
}

func main() {
	a := [][]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	b := [][]int{
		{1, 0},
		{0, 1},
		{0, 0},
	}
	if c, err := MultiplySquareMatrix(a, b); err == nil {
		fmt.Println(c)
	} else {
		fmt.Println(err)
	}
}
