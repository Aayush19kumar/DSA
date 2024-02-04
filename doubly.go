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
		fmt.Printf("%d -> ", currentNode.Data)
		currentNode = currentNode.Next
	}
	fmt.Println("nil")
}

func (dll *DoublyList) PrintReverse() {
	currentNode := dll.Tail
	for currentNode != nil {
		fmt.Printf("%d -> ", currentNode.Data)
		currentNode = currentNode.Prev
	}
	fmt.Println("nil")
}

func (dll *DoublyList) DeleteNode(node *Node) {
	if node.Next == nil || node.Prev == nil {
		fmt.Println("Next or Prev is nil")
		return
	}
	prev := node.Prev
	next := node.Next
	prev.Next = next
	next.Prev = prev
}

func main() {
	doublylist := DoublyList{}
	doublylist.AddNode(1)
	doublylist.AddNode(2)
	doublylist.AddNode(3)
	doublylist.AddNode(4)
	doublylist.AddNode(5)
	doublylist.PrintForward()
	// doublylist.PrintReverse()
	doublylist.DeleteNode(&Node{Data: 2, Next: doublylist.Head.Next.Next.Next, Prev: doublylist.Head.Next})
	doublylist.PrintForward()
	doublylist.PrintReverse()

}
