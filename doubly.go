package main

import "fmt"

type Node struct {
	Data int
	Next *Node
	Prev *Node
}
type DoublyList struct {
	Head *Node
	Tail *Node
}

func (dll *DoublyList) AddNode(data int) {
	newNode := &Node{
		Data: data,
		Next: nil,
		Prev: nil,
	}
	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		newNode.Prev = dll.Tail
		dll.Tail.Next = newNode
		dll.Tail = newNode
	}
}
func (dll *DoublyList) PrintForward() {
	currentNode := dll.Head
	for currentNode != nil {
		fmt.Printf("%d ", currentNode.Data)
		currentNode = currentNode.Next
	}
	fmt.Println()
}

func (dll *DoublyList) PrintReverse() {
	currentNode := dll.Tail
	for currentNode != nil {
		fmt.Printf("%d ", currentNode.Data)
		currentNode = currentNode.Prev
	}
	fmt.Println()
}

func main() {
	doublylist := DoublyList{}
	doublylist.AddNode(1)
	doublylist.AddNode(2)
	doublylist.AddNode(3)
	doublylist.AddNode(4)
	doublylist.AddNode(5)
	doublylist.PrintForward()
	doublylist.PrintReverse()
}
