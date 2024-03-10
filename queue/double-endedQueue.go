package main

import "fmt"

func main() {
	q := Deque{}
	fmt.Println(q.Items)
	q.PushFront(1)
	fmt.Println(q.Items)
	q.PushBack(2)
	fmt.Println(q.Items)
	q.PushFront(0)
	fmt.Println(q.Items)
	q.PopFront()
	fmt.Println(q.Items)
	q.PopBack()
	fmt.Println(q.Items)

}

type Deque struct {
	Items []int
}

func (q *Deque) PushFront(data int) {
	q.Items = append([]int{data}, q.Items...)
}
func (q *Deque) PushBack(data int) {
	q.Items = append(q.Items, data)
}

func (q *Deque) PopFront() (int, bool) {
	if len(q.Items) == 0 {
		return 0, false
	}
	front := q.Items[0]
	q.Items = q.Items[1:]
	return front, true
}

func (q *Deque) PopBack() (int, bool) {
	if len(q.Items) == 0 {
		return 0, false
	}
	rear := q.Items[len(q.Items)-1]
	q.Items = q.Items[:len(q.Items)-1]
	return rear, true
}
