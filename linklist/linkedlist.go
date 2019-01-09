package main

import "fmt"

type Element interface {
	Equal(el Element) bool
}

type Int int

func (lint Int) Equal(el Element) bool {
	if el2, ok := el.(Int); ok {
		return lint == el2 
	}
	return false
}

type Node struct {
	Key Element
	Pre *Node
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (list *LinkedList) ListSearch(k Element) *Node {
	x := list.Head;
	for x != nil && !x.Key.Equal(k) {
		x = x.Next
	}
	return x
}

func (list *LinkedList) ListInsert(x Element) {
	node := Node{x, nil, list.Head}

	if list.Head != nil {
		list.Head.Pre = &node;
	}
	list.Head = &node
}

func (list *LinkedList) ListDelete(x Element) Element {
	node := list.ListSearch(x)
	if node != nil {
		(*node).Pre.Next = (*node).Next
		(*node).Next.Pre = (*node).Pre
		(*node).Pre = nil
		(*node).Next = nil
	}

	return (*node).Key;
}

func NewLinkedList() LinkedList {
	return LinkedList{ nil, nil }
}

func SearchAndPrintResult(list LinkedList, el Element) {
	if result := list.ListSearch(el); result != nil {
		fmt.Println("found")
	} else {
		fmt.Println("not found")
	}
}

func main() {
	list := NewLinkedList()
	for i := 0; i < 100; i++ {
		list.ListInsert(Int(i))
	}

	SearchAndPrintResult(list, Int(55))
	list.ListDelete(Int(55))
	SearchAndPrintResult(list, Int(55))
}