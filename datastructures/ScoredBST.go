package datastructures

import "math"

type Scorable interface {
	Score() int     // O(1)
	String() string // O(1)
}

type Bst interface {
	Insert(Scorable)        // O(log(N))
	Remove(Scorable) Bst    // O(log(N)), returns new root
	Size() int              // O(N)
	MaxDepth() int          // O(log(N))
	MinDepth() int          // O(log(N))
	Min() Scorable          // O(log(N))
	Max() Scorable          // O(log(N))
	Contains(Scorable) bool // O(log(N))
	List() []Scorable       // O(N)
	String() string         // O(N)
}

type BstNode struct {
	value       Scorable
	score       int
	left, right *BstNode
}

func (n *BstNode) Insert(v Scorable) {
	n.insert(v, v.Score())
}

func (n *BstNode) insert(v Scorable, s int) {
	if n.score > s {
		if n.left == nil {
			n.left = &BstNode{v, s, nil, nil}
		} else {
			n.left.insert(v, s)
		}
	} else {
		if n.right == nil {
			n.right = &BstNode{v, s, nil, nil}
		} else {
			n.right.insert(v, s)
		}
	}
}

func dummyRootRemove(n *BstNode) (o *BstNode) {
	if n.left == nil {
		return n.right
	}
	if n.right == nil {
		return n.left
	}
	// Dummy root removal. Find the min value of the right subtree and make that the
	// root of this tree
	nv := n.right.Min()
	n.value = nv
	n.right = n.right.Remove(nv).(*BstNode)
	return n
}

func (n *BstNode) Remove(v Scorable) (root Bst) {
	return n.remove(v, v.Score())
}

func (n *BstNode) remove(v Scorable, s int) (root Bst) {
	if n.value == v {
		root = dummyRootRemove(n)
	} else {
		root = n
	}
	if n.score > s {
		if n.left == nil {
			return
		} else if n.left.value == v {
			n.left = dummyRootRemove(n.left)
		} else {
			n.left.remove(v, s)
		}
	} else {
		if n.right == nil {
			return
		} else if n.right.value == v {
			n.right = dummyRootRemove(n.right)
		} else {
			n.right.remove(v, s)
		}
	}
	return
}

func (n *BstNode) Min() Scorable {
	if n.left != nil {
		return n.left.Min()
	}
	return n.value
}

func (n *BstNode) Max() Scorable {
	if n.right != nil {
		return n.right.Max()
	}
	return n.value
}

func (n *BstNode) Size() int {
	size := 1
	if n.left != nil {
		size += n.left.Size()
	}
	if n.right != nil {
		size += n.right.Size()
	}
	return size
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func (n *BstNode) MaxDepth() int {
	lDepth, rDepth := 0, 0
	if n.left != nil {
		lDepth = n.left.MaxDepth()
	}
	if n.right != nil {
		rDepth = n.right.MaxDepth()
	}
	return max(lDepth, rDepth) + 1
}

func (n *BstNode) MinDepth() int {
	lDepth, rDepth := math.MaxInt32, math.MaxInt32
	if n.left != nil {
		lDepth = n.left.MaxDepth()
	}
	if n.right != nil {
		rDepth = n.right.MaxDepth()
	}
	return min(lDepth, rDepth) + 1
}

func (n *BstNode) Contains(v Scorable) (c bool) {
	if n.value == v {
		return true
	}
	if n.left != nil {
		c = c || n.left.Contains(v)
	}
	if n.right != nil {
		c = c || n.right.Contains(v)
	}
	return c
}

func (n *BstNode) List() []Scorable {
	l := make([]Scorable, 0, n.Size())
	return n.list(l)
}

func (n *BstNode) list(l []Scorable) []Scorable {
	if n.left != nil {
		l = n.left.list(l)
	}
	l = append(l, n.value)
	if n.right != nil {
		l = n.right.list(l)
	}
	return l
}

func (n *BstNode) String() string {
	s := "(" + n.value.String()
	if n.left != nil {
		s += " " + n.left.String()
	} else {
		s += " _"
	}
	if n.right != nil {
		s += " " + n.right.String()
	} else {
		s += " _"
	}
	s += ")"
	return s
}
