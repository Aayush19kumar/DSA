package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Len  int
	Head *Node
}

func (l *LinkedList) Insert(data int) {
	newNode := &Node{Data: data}

	if l.Head == nil {
		l.Head = newNode
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	l.Len++
}

func (l *LinkedList) AddFront(data int) {
	newNode := &Node{Data: data}
	if l.Head == nil {
		l.Head = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
	l.Len++
}

func (l *LinkedList) AddBack(data int) {
	newNode := &Node{Data: data}

	if l.Head == nil {
		l.Head = newNode
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	l.Len++
}

func (l *LinkedList) RemoveFront() {
	if l.Head == nil {
		fmt.Println("List is Empty")
		return
	}
	l.Head = l.Head.Next
	l.Len--

}

func (l *LinkedList) RemoveBack() {
	if l.Head == nil {
		fmt.Println("Empty List")
		return
	}
	var prev *Node
	current := l.Head
	for current.Next != nil {
		prev = current
		current = current.Next
	}
	if prev != nil {
		prev.Next = nil
	} else {
		l.Head = nil
	}
	l.Len--
}

func (l *LinkedList) Size() int {
	return l.Len
}

func (l *LinkedList) Front() int {
	if l.Head == nil {
		return 0
	}

	return l.Head.Data

}

func (l *LinkedList) Display() {

	current := l.Head
	if current == nil {
		fmt.Println("empty linked list")
		return
	}
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	list := LinkedList{}
	list.Insert(1)
	list.Insert(2)
	// list.AddFront(0)
	// list.AddFront(-1)
	// list.AddBack(3)
	// list.AddBack(4)
	// list.RemoveFront()
	// list.RemoveBack()
	// list.RemoveBack()
	// list.RemoveBack()
	list.Display()
	fmt.Println("Front", list.Front())
	fmt.Println("Length is:", list.Size())
}
