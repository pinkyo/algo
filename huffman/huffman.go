package main

import (
	"container/heap"
	"fmt"
)

type FreData struct {
	char  string
	fre   float64
	left  *FreData
	right *FreData
}

type FreDataHeap []FreData

func (h FreDataHeap) Len() int           { return len(h) }
func (h FreDataHeap) Less(i, j int) bool { return h[i].fre < h[j].fre }
func (h FreDataHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *FreDataHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(FreData))
}

func (h *FreDataHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func huffnanCode(fre FreDataHeap) FreData {
	// code := make([]string, len(fre))
	heap.Init(&fre)
	for fre.Len() > 1 {
		left := heap.Pop(&fre).(FreData)
		right := heap.Pop(&fre).(FreData)
		heap.Push(&fre, FreData{"", right.fre + left.fre, &left, &right})
	}

	root := heap.Pop(&fre).(FreData)

	return root
}

func printlnTree(root FreData, code string) {
	if root.left == nil && root.right == nil {
		fmt.Println(root.char, ":", code)
		return
	}

	if root.left != nil {
		printlnTree(*root.left, code+"0")
	}
	if root.right != nil {
		printlnTree(*root.right, code+"1")
	}
}

func main() {
	charFreArr := make([]FreData, 6)
	charFreArr[0] = FreData{"f", 5, nil, nil}
	charFreArr[1] = FreData{"e", 9, nil, nil}
	charFreArr[2] = FreData{"c", 12, nil, nil}
	charFreArr[3] = FreData{"b", 13, nil, nil}
	charFreArr[4] = FreData{"d", 16, nil, nil}
	charFreArr[5] = FreData{"a", 45, nil, nil}

	root := huffnanCode(charFreArr)
	printlnTree(root, "")
}
