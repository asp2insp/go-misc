package datastructures

import "fmt"

type T interface{}

type CircularQueue interface {
	Enqueue(T)           // O(1)
	Dequeue() (T, error) //O(1)
	Size() int           // O(1)
}

type ArrayCircularQueue struct {
	data             []T
	size, start, end int
}

func CreateQueue() (ret *ArrayCircularQueue) {
	ret = new(ArrayCircularQueue)
	ret.data = make([]T, 4)
	return
}

func (a *ArrayCircularQueue) Size() int {
	return a.size
}

func (a *ArrayCircularQueue) reallocate() {
	nq := make([]T, a.size*2)
	for i := 0; i < a.size; i++ {
		nq[i] = a.data[(a.start+i)%len(a.data)]
	}
	a.data = nq
	a.start = 0
}

func (a *ArrayCircularQueue) Enqueue(val T) {
	if a.size == len(a.data) {
		a.reallocate()
	}
	a.data[(a.start+a.size)%len(a.data)] = val
	a.size++
}

func (a *ArrayCircularQueue) Dequeue() (val T, err error) {
	if a.size == 0 {
		err = fmt.Errorf("Queue is empty")
	} else {
		val = a.data[a.start]
		a.start++
		a.size--
	}
	return
}

func (a *ArrayCircularQueue) String() string {
	s := "["
	for i := 0; i < a.size; i++ {
		if i != 0 {
			s += ","
		}
		s += a.data[(a.start+i)%len(a.data)].(string)
	}
	s += "]"
	return s
}
