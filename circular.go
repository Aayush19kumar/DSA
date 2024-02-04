package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}
type CircularList struct {
	Head *Node
	Tail *Node
}

func (cl *CircularList) Insert(data int) {
	newNode := &Node{Data: data, Next: nil}

	if cl.Head == nil {
		cl.Head = newNode
		cl.Tail = newNode
	} else {
		cl.Tail.Next = newNode
		cl.Tail = newNode
		newNode.Next = cl.Head
	}

}

func (cl *CircularList) Traverse() {
	current := cl.Head
	for {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
		if current == cl.Head {
			break
		}
	}
	fmt.Println("nil")
}

func main() {
	cll := CircularList{}
	cll.Insert(1)
	cll.Insert(2)
	cll.Insert(3)
	cll.Insert(4)
	cll.Insert(5)
	cll.Traverse()
}
