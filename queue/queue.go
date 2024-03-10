package main

import "fmt"

func main() {
	queue := Queue{Size: 3}
	fmt.Println(queue.Element)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println(queue.Element)
	queue.Deque()
	queue.Deque()
	queue.Deque()
	fmt.Println(queue.Element)
	queue.Deque()

}

type Queue struct {
	Element []int
	Size    int
}

func (q *Queue) GetLength() int {
	return len(q.Element)
}
func (q *Queue) IsEmpty() bool {
	return len(q.Element) == 0
}
func (q *Queue) Enqueue(data int) {
	if q.GetLength() == q.Size {
		fmt.Println("OverFlow")
		return
	}
	q.Element = append(q.Element, data)
}
func (q *Queue) Deque() int {
	if q.IsEmpty() {
		fmt.Println("UnderFlow")
		return 0
	}
	res := q.Element[0]
	if q.GetLength() == 1 {
		q.Element = nil
		return res
	}
	q.Element = q.Element[1:]
	return res
}

func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("empty queue")
	}
	return q.Element[0], nil
}
