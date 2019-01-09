package main

import (
	_ "errors"
	"fmt"
)

type Element interface {
	Compare(el Element) int
}

type RBColor int

const (
	RED RBColor = iota
	BLACK
)

type Node struct {
	Key   Element
	Left  *Node
	Right *Node
	P     *Node
	Color RBColor
}

func TreeSearch(x *Node, k Element) *Node {
	node := x
	for node != nil && node.Key.Compare(k) != 0 {
		if node.Key.Compare(k) > 0 {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	return node
}

func TreeInorderWalk(x *Node) {
	if x == nil {
		return
	}
	if x.Left != nil {
		TreeInorderWalk(x.Left)
	}
	fmt.Print(x.Key, " ")
	if x.Right != nil {
		TreeInorderWalk(x.Right)
	}
}

func TreeMinimun(node *Node) *Node {
	for node.Left != nil {
		node = node.Left
	}

	return node
}

func TreeMaximun(x *Node) *Node {
	node := x
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func TreeSuccessor(x *Node) *Node {
	node := x
	if node.Right != nil {
		return TreeMinimun(node.Right)
	} else {
		y := node.P
		for y != nil && y.Right == node {
			x = y
			y = y.P
		}

		return y
	}
}

func TreePrecessor(x *Node) *Node {
	node := x
	if node.Left != nil {
		return TreeMaximun(node.Left)
	} else {
		y := node
		for y != nil && y.Left == node {
			node = y
			y = y.P
		}

		return y
	}
}

type Bst struct {
	Root *Node
}

func (bst *Bst) TreeInsert(z Element) {
	var y (*Node)
	y = nil
	x := bst.Root
	for x != nil {
		y = x
		if x.Key.Compare(z) < 0 {
			x = x.Right
		} else {
			x = x.Left
		}
	}

	node := new(Node)
	node.Key = z
	node.P = y
	node.Left = nil
	node.Right = nil

	if y == nil {
		bst.Root = node
	} else if node.Key.Compare(y.Key) < 0 {
		y.Left = node
	} else {
		y.Right = node
	}
}

func (bst *Bst) Transplant(u *Node, v *Node) {
	if u.P == nil {
		bst.Root = v
	} else if u == u.P.Left {
		u.P.Left = v
	} else {
		u.P.Right = v
	}
	if v != nil {
		v.P = u.P
	}
}

func (bst *Bst) TreeDelete(z *Node) {
	if z.Left == nil {
		bst.Transplant(z, z.Right)
	} else if z.Right == nil {
		bst.Transplant(z, z.Left)
	} else {
		y := TreeMinimun(z.Right)
		if y.P != z {
			bst.Transplant(y, y.Right)
			y.Right = z.Right
			y.Right.P = z.P
		}

		bst.Transplant(z, y)
		y.Left = z.Left
		y.Left.P = y
	}
}

func (bst *Bst) LeftRotate(x *Node) {
	y := x.Right
	x.Right = y.Left
	if y.Left != nil {
		y.Left.P = x
	}
	y.P = x.P
	if x.P == nil {
		bst.Root = y
	} else if x == x.P.Left {
		x.P.Left = y
	} else {
		x.P.Right = y
	}
	y.Left = x
	x.P = y
}

func (bst *Bst) RightRotate(y *Node) {
	x := y.Left
	y.Left = x.Right
	if x.Right != nil {
		x.Right.P = y
	}
	x.P = y.P
	if y.P == nil {
		bst.Root = x
	} else if y == x.P.Left {
		y.P.Left = x
	} else {
		y.P.Right = x
	}
	x.Left = y
	y.P = x
}

func (bst *Bst) RbInsert(z *Node) {
	var y *Node
	x := bst.Root

	for x != nil {
		y = x
		if z.Key.Compare(x.Key) < 0 {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	z.P = y
	if y == nil {
		bst.Root = z
	} else if z.Key.Compare(y.Key) < 0 {
		y.Left = z
	} else {
		y.Right = z
	}

	z.Left = nil
	z.Right = nil
	z.Color = RED
	bst.RbInsertFixup(z)
}

func (bst *Bst) RbInsertFixup(z *Node) {
	for z.P.Color == RED {
		if z.P == z.P.P.Left {
			y := z.P.P.Right
			if y.Color == RED {
				z.P.Color = BLACK
				y.Color = BLACK
				z.P.P.Color = RED
				z = z.P.P
			} else if z == z.P.Right {
				z = z.P
				bst.LeftRotate(z)
			} else {
				z.P.Color = BLACK
				z.P.P.Color = RED
				bst.RightRotate(z.P.P)
			}
		} else {
			y := z.P.P.Left
			if y.Color == RED {
				z.P.Color = BLACK
				y.Color = BLACK
				z.P.P.Color = RED
				z = z.P.P
			} else if z == z.P.Left {
				z = z.P
				bst.RightRotate(z)
			} else {
				z.P.Color = BLACK
				z.P.P.Color = RED
				bst.LeftRotate(z.P.P)
			}
		}
	}

	bst.Root.Color = BLACK
}

func (bst *Bst) RbTransplant(u *Node, v *Node) {
	if u.P == nil {
		bst.Root = v
	} else if u == u.P.Left {
		u.P.Left = v
	} else {
		u.P.Right = v
	}
	if v != nil {
		v.P = u.P
	}
}

func (bst *Bst) RbDelete(z *Node) {
	y := z
	yOriginalColor := y.Color
	var x *Node
	if z.Left == nil {
		x = z.Right
		bst.RbTransplant(z, z.Right)
	} else if z.Right == nil {
		x = z.Left
		bst.RbTransplant(z, z.Left)
	} else {
		y = TreeMinimun(z.Right)
		yOriginalColor = y.Color
		x = y.Right
		if y.P == z && x != nil {
			x.P = y
		} else {
			bst.RbTransplant(y, y.Right)
		}
		bst.RbTransplant(z, y)
		y.Left = z.Left
		y.Left.P = y
		y.Color = z.Color
	}

	if yOriginalColor == BLACK {
		bst.RbDeleteFixup(x)
	}
}

func (bst *Bst) RbDeleteFixup(x *Node) {
	for x != bst.Root && x.Color == BLACK {
		if x == x.P.Left {
			w := x.P.Right
			if w.Color == RED {
				w.Color = BLACK
				x.P.Color = RED
				bst.LeftRotate(x.P)
				w = x.P.Right
			}

			if w.Left.Color == BLACK && w.Right.Color == BLACK {
				w.Color = RED
				x = x.P
			} else if w.Right.Color == BLACK {
				w.Left.Color = BLACK
				w.Color = RED
				bst.RightRotate(w)
				w = x.P.Right
			} else {
				w.Color = x.P.Color
				x.P.Color = BLACK
				w.Right.Color = BLACK
				bst.LeftRotate(x.P)
				x = bst.Root
			}
		} else {
			w := x.P.Left
			if w.Color == RED {
				w.Color = BLACK
				x.P.Color = RED
				bst.RightRotate(x.P)
				w = x.P.Left
			}

			if w.Left.Color == BLACK && w.Right.Color == BLACK {
				w.Color = RED
				x = x.P
			} else if w.Left.Color == BLACK {
				w.Right.Color = BLACK
				w.Color = RED
				bst.LeftRotate(w)
				w = x.P.Left
			} else {
				w.Color = x.P.Color
				x.P.Color = BLACK
				w.Left.Color = BLACK
				bst.RightRotate(x.P)
				x = bst.Root
			}
		}
	}

	x.Color = BLACK
}

type Int int

func (lint Int) Compare(el Element) int {
	if i, ok := el.(Int); ok {
		return int(lint) - int(i)
	}
	return -1
}

func NewBst() Bst {
	return Bst{nil}
}

func main() {
	bst := NewBst()
	for i := 0; i < 10; i++ {
		bst.TreeInsert(Int(i))
	}
	TreeInorderWalk(bst.Root)
	fmt.Println()

	if res := TreeSearch(bst.Root, Int(55)); res != nil {
		fmt.Println("found")
	} else {
		fmt.Println("not found")
	}

	var node *Node
	for i := 0; i < 10; i++ {
		if node = TreeSearch(bst.Root, Int(i)); node != nil {
			bst.TreeDelete(node)
		}
	}
	TreeInorderWalk(bst.Root)
	fmt.Println()

}
