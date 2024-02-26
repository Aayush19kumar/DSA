spackage main

import (
	"fmt"
)

type (
	Node struct {
		Data int
		Next *Node
	}

	LinkedList struct {
		Len  int
		Head *Node
	}
)

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

func (l *LinkedList) Middle() *Node {
	if l.Head == nil {
		fmt.Println("[Middle]Empty list")
		return nil
	}
	fast := l.Head
	slow := l.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

	}
	return slow
}

func (l *LinkedList) InsertAtASpecificValue(k int) {
	newNode := &Node{Data: k + 1}
	if l.Head == nil {
		fmt.Println("[InsertAtASpecificValue] Empty List")
		return
	}
	current := l.Head
	for current.Data != k {
		current = current.Next
	}

	temp := current.Next
	current.Next = newNode
	newNode.Next = temp
}

func (l *LinkedList) DeleteAtASpecificValue(k int) *Node {
	if l.Head == nil {
		fmt.Println("[DeleteAtASpecificValue] Empty List")
		return nil
	}

	if l.Head.Data == k {
		return l.Head.Next
	}

	current := l.Head
	prev := l.Head

	for current.Data != k && current != nil {
		prev = current
		current = current.Next
	}
	if current != nil {
		prev.Next = current.Next
	}
	return l.Head
}

func (l *LinkedList) Display() {

	current := l.Head
	if current == nil {
		fmt.Println("[Display] empty linked list")
		return
	}
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}
func Show(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}
func (l *LinkedList) Reverse() *Node {

	var current, prev, next = l.Head, *Node, *Node
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next

	}
	return prev
}

func printReverseOrder(node *Node) {
	if node != nil {
		printReverseOrder(node.Next)
		fmt.Printf("%d ", node.Data)
	}
}

func swap(node *Node) {
	temp := node.Data
	node.Data = node.Next.Data
	node.Next.Data = temp

	current := node
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")

}

func DeleteParticularNode(node *Node) {
	current := node
	prev := node
	for current != nil && current.Next != nil {
		current.Data, current.Next.Data = current.Next.Data, current.Data
		prev = current
		current = current.Next
	}
	prev.Next = nil

	for node != nil {
		fmt.Printf("%d -> ", node.Data)
		node = node.Next
	}
	fmt.Println("nil")

}

func reverseList(head *Node) *Node {
    current:=head
	Tail:=head
	for Tail!=nil{
		Tail=Tail.next
	}

}

func main() {
	list := LinkedList{}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	// list.Insert(5)
	// list.AddFront(0)
	// list.AddFront(-1)
	// list.AddBack(3)
	// list.AddBack(4)
	// list.RemoveFront()
	// list.RemoveBack()
	// list.RemoveBack()
	// list.RemoveBack()
	list.Display()
	// fmt.Println("Front", list.Front())
	// fmt.Println("Length is:", list.Size())
	// middle := list.Middle()
	// if middle != nil {
	// 	fmt.Println("Middle of the list is :", list.Middle().Data)
	// } else {
	// 	fmt.Println("Middle of the list is:", 0)
	// }
	// list.InsertAtASpecificValue(2)
	// list.Display()
	// Show(list.DeleteAtASpecificValue(3))
	// Show(list.Reverse()
	// swap(list.Head.Next)
	DeleteParticularNode(list.Head.Next)

}
