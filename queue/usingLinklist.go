package main

import (
	"container/list"
	"fmt"
)

func main() {
	// new linked list
	queue := list.New()
	queue.PushBack(10)
	queue.PushBack(20)
	queue.PushBack(30)

	// Deque
	front := queue.Front()
	fmt.Println(front.Value)
	queue.Remove(front)
}
