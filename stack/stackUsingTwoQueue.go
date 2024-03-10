package main

import "fmt"

func main() {

	s := NewStack()
	s.Push(1)
	s.Push(2)
	fmt.Println(s.Peek())

}

type Stackk struct {
	q1 []interface{}
	q2 []interface{}
}

func NewStack() *Stackk {
	return &Stackk{
		q1: make([]interface{}, 0),
		q2: make([]interface{}, 0),
	}
}

func (q *Stackk) Push(data int) {
	q.q2 = append(q.q2, data)

	for (len(q.q1)) > 0 {
		elem := q.q1[len(q.q1)-1]
		q.q1 = q.q1[:len(q.q1)-1]
		q.q2 = append(q.q2, elem)
	}
	q.q1, q.q2 = q.q2, q.q1
}

func (q *Stackk) PoP() interface{} {
	if len(q.q1) == 0 {
		return nil
	}
	elem := q.q1[len(q.q1)-1]
	q.q1 = q.q1[:len(q.q1)-1]
	return elem
}

func (q *Stackk) Peek() interface{} {
	if len(q.q1) == 0 {
		return nil
	}
	return q.q1[len(q.q1)-1]

}
