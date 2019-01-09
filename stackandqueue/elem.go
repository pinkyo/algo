package main

import (
	"fmt"
	"errors"
)
type Element interface {

}

type Stack struct {
	data []Element
	top int
}

func (stack *Stack) IsEmpty() bool {
	return stack.top == 0
}

func Resize(oldData []Element) []Element {
	newData := make([]Element, len(oldData) * 2 + 1)
	return newData
}

func (stack *Stack) Push(el Element) {
	if len(stack.data) == stack.top {
		newData := Resize(stack.data)
		copy(newData, stack.data)
		stack.data = newData
	}

	stack.data[stack.top] = el
	stack.top ++
}

func (stack *Stack) Pop() (Element, error) {
	if stack.top == 0 {
		return nil, errors.New("underflow")
	}
	
	stack.top --
	return stack.data[stack.top], nil
}

func NewStack() Stack {
	return Stack{make([] Element, 10), 0}
}

type Queue struct {
	data []Element
	head int
	tail int
}

func (queue *Queue) Enqueue(el Element) {
	if (queue.tail + 1) % len(queue.data) == queue.head {
		newData := Resize(queue.data)

		i := queue.head
		for j := 0; j < len(queue.data) - 1; j++ {
			newData[j] = queue.data[i]
			i = (i + 1) % len(queue.data)
		}
		queue.head, queue.tail = 0, len(queue.data) - 1
		queue.data = newData
	}

	queue.data[queue.tail] = el
	queue.tail = (queue.tail + 1) % len(queue.data)
}

func (queue *Queue) Dequeue() (Element, error) {
	if queue.head == queue.tail {
		return nil, errors.New("underflow")
	}
	result := queue.data[queue.head]
	queue.head = (queue.head + 1) % len(queue.data)

	return result, nil
}

func (queue *Queue) Size() int {
	return (queue.tail + len(queue.data) - queue.head) % len(queue.data)
}

func NewQueue() Queue {
	return Queue{make([]Element, 10), 0, 0}
}

func main() {
	stack := NewStack()
	fmt.Println(stack.IsEmpty())
	for i := 1; i < 10000; i++ {
		stack.Push(i)
	}
	fmt.Println(stack.IsEmpty())

	for i := 9999; i >= 1; i-- {
		el, _ := stack.Pop()
		if el != i {
			fmt.Println(i, el, "STACK FAIL.")
		}
	}

	queue := NewQueue()
	for i := 1; i < 10000; i++ {
		queue.Enqueue(i)
	}

	fmt.Println(queue.Size())

	for i := 1; i < 10000; i++ {
		el, _ := queue.Dequeue()
		if el != i {
			fmt.Println(i, el, "QUEUE FAIL.")
		}
	} 
}