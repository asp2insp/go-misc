package datastructures

import "strconv"
import "fmt"

type Node struct {
	value      int
	prev, next *Node
}

type IntegerLinkedList interface {
	Add(int)                      // O(1)
	Remove(int)                   // O(N)
	RemoveIndex(int) (int, error) // O(N)
	Get(int) (int, error)         // O(N)
	IsEmpty() bool                // O(1)
	Contains(int) bool            // O(N)
	PushBack(int)                 // O(1)
	PushFront(int)                // O(1)
	PopBack() (int, error)        // O(1)
	PopFront() (int, error)       // O(1)
	Size() int                    // O(1)
}

type DoublyLinkedIntegerList struct {
	head, tail *Node
	size       int
}

func (dll *DoublyLinkedIntegerList) init(n *Node) {
	dll.head = n
	dll.tail = n
}

func (dll *DoublyLinkedIntegerList) String() string {
	var s string
	s += "["
	n := dll.head
	for n != nil {
		s += strconv.Itoa(n.value)
		s += ","
		n = n.next
	}
	s += "]"
	return s
}

func (dll *DoublyLinkedIntegerList) Contains(val int) bool {
	n := dll.head
	for n != nil {
		if n.value == val {
			return true
		}
		n = n.next
	}
	return false
}

func (dll *DoublyLinkedIntegerList) spliceOut(n *Node) {
	if n == nil {
		return
	}
	if n.next != nil {
		n.next.prev = n.prev
	} else {
		// n is tail, need to update
		dll.tail = n.prev
	}
	if n.prev != nil {
		n.prev.next = n.next
	} else {
		// n is head, need to update
		dll.head = n.next
	}
	dll.size -= 1
}

func (dll *DoublyLinkedIntegerList) Remove(val int) {
	n := dll.head
	for n != nil {
		if n.value == val {
			dll.spliceOut(n)
			return
		}
		n = n.next
	}
}

func (dll *DoublyLinkedIntegerList) getNode(index int) (node *Node, err error) {
	if index >= dll.size {
		err = fmt.Errorf("Index %d out of bounds", index)
	} else if index < 0 {
		n := dll.tail
		for ; index < -1; index++ {
			n = n.prev
		}
		node = n
		err = nil
	} else {
		n := dll.head
		for ; index > 0; index-- {
			n = n.next
		}
		node = n
		err = nil
	}
	return
}

func (dll *DoublyLinkedIntegerList) Get(index int) (value int, err error) {
	n, err := dll.getNode(index)
	if n != nil {
		value = n.value
	}
	return
}

func (dll *DoublyLinkedIntegerList) RemoveIndex(index int) (value int, err error) {
	n, err := dll.getNode(index)
	if n != nil {
		value = n.value
		dll.spliceOut(n)
	}
	return
}

func (dll *DoublyLinkedIntegerList) Size() int {
	return dll.size
}

func (dll *DoublyLinkedIntegerList) IsEmpty() bool {
	return (dll.size == 0)
}

func (dll *DoublyLinkedIntegerList) Add(n int) {
	dll.PushBack(n)
}

func (dll *DoublyLinkedIntegerList) PushBack(n int) {
	node := new(Node)
	node.value = n
	if dll.tail == nil {
		dll.init(node)
	} else {
		dll.tail.next = node
		node.prev = dll.tail
		dll.tail = node
	}
	dll.size += 1
}

func (dll *DoublyLinkedIntegerList) PushFront(n int) {
	node := new(Node)
	node.value = n
	if dll.head == nil {
		dll.init(node)
	} else {
		node.next = dll.head
		dll.head.prev = node
		dll.head = node
	}
	dll.size += 1
}

func (dll *DoublyLinkedIntegerList) PopBack() (value int, err error) {
	if dll.tail == nil {
		err = fmt.Errorf("PopBack called on empty list")
	} else {
		value = dll.tail.value
		dll.spliceOut(dll.tail)
	}
	return
}

func (dll *DoublyLinkedIntegerList) PopFront() (value int, err error) {
	if dll.head == nil {
		err = fmt.Errorf("PopFront called on empty list")
	} else {
		value = dll.head.value
		dll.spliceOut(dll.head)
	}
	return
}
