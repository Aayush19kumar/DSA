// package main

// import "fmt"

// type Node struct {
// 	Data int
// 	Next *Node
// 	Prev *Node
// }

// type DoublyLinkedList struct {
// 	Head *Node
// 	Tail *Node
// }

// func (dll *DoublyLinkedList) Insert(Data int) {
// 	newNode := &Node{Data: Data, Next: nil, Prev: nil}

// 	if dll.Head == nil {
// 		dll.Head = newNode
// 		dll.Tail = newNode
// 	} else {
// 		newNode.Prev = dll.Tail
// 		dll.Tail.Next = newNode
// 		dll.Tail = newNode
// 	}

// }
// func (dll *DoublyLinkedList) PrintForward() {
// 	current := dll.Head
// 	for current != nil {
// 		fmt.Printf("%d -> ", current.Data)
// 		current = current.Next
// 	}
// 	fmt.Println("nil")
// }
// func (dll *DoublyLinkedList) PrintREverse() {
// 	current := dll.Tail
// 	for current != nil {
// 		fmt.Printf("%d -> ", current.Data)
// 		current = current.Prev
// 	}
// 	fmt.Println("nil")
// }

// func main() {
// 	dll := DoublyLinkedList{}
// 	dll.Insert(1)
// 	dll.Insert(2)
// 	dll.Insert(3)
// 	dll.Insert(4)
// 	dll.Insert(5)
// 	dll.PrintForward()
// 	dll.PrintREverse()

// }
