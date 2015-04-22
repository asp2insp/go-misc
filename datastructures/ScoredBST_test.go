package datastructures

import "strconv"
import "testing"
import "github.com/asp2insp/go-misc/testutils"

type IntWrap struct {
	int
}

func (i IntWrap) Score() int {
	return i.int
}

func (i IntWrap) String() string {
	return strconv.Itoa(i.int)
}

func TestSatisfies(t *testing.T) {
	var bst Bst
	a := &BstNode{IntWrap{3}, 3, nil, nil}
	bst = a
	testutils.CheckInt(1, bst.Size(), t)
}

func TestTreeInsert(t *testing.T) {
	n := &BstNode{IntWrap{3}, 3, nil, nil}
	n.Insert(IntWrap{6})
	n.Insert(IntWrap{2})
	n.Insert(IntWrap{9})

	testutils.ExpectTrue(n.Contains(IntWrap{6}), "Should contain 6", t)
	testutils.ExpectTrue(n.Contains(IntWrap{3}), "Should containg 3", t)
	testutils.ExpectTrue(n.Contains(IntWrap{2}), "Should contain 2", t)
	testutils.ExpectTrue(n.Contains(IntWrap{9}), "Should contain 9", t)

	testutils.CheckInt(4, n.Size(), t)

	testutils.CheckString("(3 (2 _ _) (6 _ (9 _ _)))", n.String(), t)
}

func TestTreeRemove(t *testing.T) {
	var n Bst = &BstNode{IntWrap{3}, 3, nil, nil}
	n.Insert(IntWrap{6})
	n.Insert(IntWrap{2})
	n.Insert(IntWrap{9})

	testutils.ExpectTrue(n.Contains(IntWrap{6}), "Should contain 6", t)
	testutils.CheckString("(3 (2 _ _) (6 _ (9 _ _)))", n.String(), t)

	n = n.Remove(IntWrap{6})
	testutils.ExpectFalse(n.Contains(IntWrap{6}), "6 was removed", t)
	testutils.CheckString("(3 (2 _ _) (9 _ _))", n.String(), t)

	n = n.Remove(IntWrap{2})
	testutils.ExpectFalse(n.Contains(IntWrap{2}), "2 was removed", t)
	testutils.CheckString("(3 _ (9 _ _))", n.String(), t)

	testutils.CheckInt(2, n.Size(), t)
}

func TestTreeRemoveRoot(t *testing.T) {
	var n Bst = &BstNode{IntWrap{3}, 3, nil, nil}
	n.Insert(IntWrap{6})
	n.Insert(IntWrap{2})
	n.Insert(IntWrap{9})

	testutils.ExpectTrue(n.Contains(IntWrap{6}), "Should contain 6", t)
	testutils.CheckString("(3 (2 _ _) (6 _ (9 _ _)))", n.String(), t)

	n = n.Remove(IntWrap{3})
	testutils.ExpectFalse(n.Contains(IntWrap{3}), "3 was removed", t)
	testutils.CheckString("(6 (2 _ _) (9 _ _))", n.String(), t)
}

func TestTreeStats(t *testing.T) {
	n := &BstNode{IntWrap{3}, 3, nil, nil}
	n.Insert(IntWrap{6})
	n.Insert(IntWrap{2})
	n.Insert(IntWrap{9})

	testutils.CheckInt(4, n.Size(), t)
	testutils.CheckString("(3 (2 _ _) (6 _ (9 _ _)))", n.String(), t)
	testutils.CheckInt(2, n.MinDepth(), t)
	testutils.CheckInt(3, n.MaxDepth(), t)
	testutils.ExpectTrue(IntWrap{9} == n.Max(), "9 should be max", t)
	testutils.ExpectTrue(IntWrap{2} == n.Min(), "2 should be min", t)
}

func TestTreeList(t *testing.T) {
	n := &BstNode{IntWrap{3}, 3, nil, nil}
	n.Insert(IntWrap{6})
	n.Insert(IntWrap{2})
	n.Insert(IntWrap{9})

	expected := []Scorable{IntWrap{2}, IntWrap{3}, IntWrap{6}, IntWrap{9}}
	checkSlice(expected, n.List(), t)
}

func checkSlice(exp, obs []Scorable, t *testing.T) {
	if len(exp) != len(obs) {
		t.Errorf("Differing slice lengths: %d vs %d", len(exp), len(obs))
	}
	for i, v := range exp {
		if v != obs[i] {
			t.Errorf("Slices differ in element number %d: %v vs %v", i, v, obs[i])
		}
	}
}
