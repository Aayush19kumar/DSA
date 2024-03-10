package main

import "fmt"

func main() {
	qu := QueueUsingStack{}
	qu.Enqueue(1)
	qu.Enqueue(2)
	qu.Enqueue(3)
	fmt.Println(qu.Dequeue())
	fmt.Println(qu.Dequeue())
	fmt.Println(qu.Dequeue())
	fmt.Println(qu.Dequeue())

}

type QueueUsingStack struct {
	InputStack   []int
	OutputtStack []int
}

func (q *QueueUsingStack) Enqueue(data int) {
	q.InputStack = append(q.InputStack, data)
}
func (q *QueueUsingStack) DequeFromInput() {
	for len(q.InputStack) > 0 {
		top := q.InputStack[len(q.InputStack)-1]
		q.InputStack = q.InputStack[:len(q.InputStack)-1]
		q.OutputtStack = append(q.OutputtStack, top)
	}
}

func (q *QueueUsingStack) Dequeue() int {
	if len(q.OutputtStack) == 0 {
		q.DequeFromInput()
	}
	if len(q.OutputtStack) > 0 {
		top := q.OutputtStack[len(q.OutputtStack)-1]
		q.OutputtStack = q.OutputtStack[:len(q.OutputtStack)-1]
		return top
	}
	return -1
}
