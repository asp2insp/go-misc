package datastructures

import "testing"

import "github.com/ca-geo/go-misc/testutils"

func checkFulfillsInterface(t *testing.T) {
	var ll IntegerLinkedList
	dll := new(DoublyLinkedIntegerList)
	ll = dll
	testutils.CheckInt(0, ll.Size(), t)
}

func TestPushBack(t *testing.T) {
	dll := new(DoublyLinkedIntegerList)
	testutils.CheckString("[]", dll.String(), t)
	testutils.CheckInt(0, dll.Size(), t)
	testutils.ExpectTrue(dll.IsEmpty(), "List should be empty", t)
	dll.PushBack(5)
	testutils.CheckString("[5,]", dll.String(), t)
	testutils.CheckInt(1, dll.Size(), t)
	testutils.ExpectFalse(dll.IsEmpty(), "List shouldn't be empty", t)
	dll.PushBack(10)
	testutils.CheckString("[5,10,]", dll.String(), t)
	testutils.CheckInt(2, dll.Size(), t)
}

func TestPushFront(t *testing.T) {
	dll := new(DoublyLinkedIntegerList)
	testutils.CheckString("[]", dll.String(), t)
	testutils.CheckInt(0, dll.Size(), t)
	testutils.ExpectTrue(dll.IsEmpty(), "List should be empty", t)
	dll.PushFront(5)
	testutils.CheckString("[5,]", dll.String(), t)
	testutils.CheckInt(1, dll.Size(), t)
	testutils.ExpectFalse(dll.IsEmpty(), "List shouldn't be empty", t)
	dll.PushFront(10)
	testutils.CheckString("[10,5,]", dll.String(), t)
	testutils.CheckInt(2, dll.Size(), t)
}

func TestContains(t *testing.T) {
	dll := new(DoublyLinkedIntegerList)
	dll.PushFront(5)
	dll.PushFront(10)
	testutils.ExpectTrue(dll.Contains(5), "List should contain 5", t)
	testutils.ExpectTrue(dll.Contains(10), "List should contain 10", t)
	testutils.ExpectFalse(dll.Contains(13), "List should not contain 13", t)
}

func TestGet(t *testing.T) {
	dll := new(DoublyLinkedIntegerList)
	dll.PushBack(5)
	dll.PushBack(10)
	dll.PushBack(36)
	dll.PushBack(42)
	val, err := dll.Get(0)
	if err == nil {
		testutils.CheckInt(5, val, t)
	} else {
		t.Error(err.Error())
	}
	val, err = dll.Get(1)
	if err == nil {
		testutils.CheckInt(10, val, t)
	} else {
		t.Error(err.Error())
	}
	val, err = dll.Get(2)
	if err == nil {
		testutils.CheckInt(36, val, t)
	} else {
		t.Error(err.Error())
	}

	// Negative indicies
	val, err = dll.Get(-1)
	if err == nil {
		testutils.CheckInt(42, val, t)
	} else {
		t.Error(err.Error())
	}
	val, err = dll.Get(-2)
	if err == nil {
		testutils.CheckInt(36, val, t)
	} else {
		t.Error(err.Error())
	}
}

func TestRemove(t *testing.T) {
	dll := new(DoublyLinkedIntegerList)
	dll.PushBack(5)
	dll.PushBack(10)
	dll.PushBack(36)
	dll.PushBack(42)
	testutils.CheckString("[5,10,36,42,]", dll.String(), t)
	testutils.CheckInt(4, dll.Size(), t)

	dll.Remove(36)
	testutils.CheckString("[5,10,42,]", dll.String(), t)
	testutils.CheckInt(3, dll.Size(), t)

	dll.Remove(45)
	testutils.CheckString("[5,10,42,]", dll.String(), t)
	testutils.CheckInt(3, dll.Size(), t)

	dll.Remove(42)
	testutils.CheckString("[5,10,]", dll.String(), t)
	testutils.CheckInt(2, dll.Size(), t)

	dll.Remove(5)
	testutils.CheckString("[10,]", dll.String(), t)
	testutils.CheckInt(1, dll.Size(), t)

	dll.Remove(10)
	testutils.CheckString("[]", dll.String(), t)
	testutils.CheckInt(0, dll.Size(), t)
	testutils.ExpectTrue(dll.IsEmpty(), "List should be empty now", t)
}

func TestRemoveIndex(t *testing.T) {
	dll := new(DoublyLinkedIntegerList)
	dll.PushBack(5)
	dll.PushBack(10)
	dll.PushBack(36)
	dll.PushBack(42)
	testutils.CheckString("[5,10,36,42,]", dll.String(), t)
	testutils.CheckInt(4, dll.Size(), t)

	val, err := dll.RemoveIndex(36)
	if err == nil {
		t.Error("Wanted error, but didn't get it")
	}
	testutils.CheckString("[5,10,36,42,]", dll.String(), t)
	testutils.CheckInt(4, dll.Size(), t)

	val, err = dll.RemoveIndex(2)
	if err != nil {
		t.Errorf("Got error %s", err.Error())
	} else {
		testutils.CheckInt(36, val, t)
	}
	testutils.CheckString("[5,10,42,]", dll.String(), t)
	testutils.CheckInt(3, dll.Size(), t)

	val, err = dll.RemoveIndex(2)
	if err != nil {
		t.Errorf("Got error %s", err.Error())
	} else {
		testutils.CheckInt(42, val, t)
	}
	testutils.CheckString("[5,10,]", dll.String(), t)
	testutils.CheckInt(2, dll.Size(), t)

	val, err = dll.RemoveIndex(0)
	if err != nil {
		t.Errorf("Got error %s", err.Error())
	} else {
		testutils.CheckInt(5, val, t)
	}
	testutils.CheckString("[10,]", dll.String(), t)
	testutils.CheckInt(1, dll.Size(), t)

	val, err = dll.RemoveIndex(0)
	if err != nil {
		t.Errorf("Got error %s", err.Error())
	} else {
		testutils.CheckInt(10, val, t)
	}
	testutils.CheckString("[]", dll.String(), t)
	testutils.CheckInt(0, dll.Size(), t)
	testutils.ExpectTrue(dll.IsEmpty(), "List should be empty now", t)
}

func TestPopBack(t *testing.T) {
	dll := new(DoublyLinkedIntegerList)
	dll.PushBack(5)
	dll.PushBack(10)
	dll.PushBack(36)
	dll.PushBack(42)
	testutils.CheckString("[5,10,36,42,]", dll.String(), t)
	testutils.CheckInt(4, dll.Size(), t)

	val, err := dll.PopBack()
	if err != nil {
		t.Errorf("Got error %s", err.Error())
	} else {
		testutils.CheckInt(42, val, t)
	}
	testutils.CheckString("[5,10,36,]", dll.String(), t)
	testutils.CheckInt(3, dll.Size(), t)

	val, err = dll.PopBack()
	if err != nil {
		t.Errorf("Got error %s", err.Error())
	} else {
		testutils.CheckInt(36, val, t)
	}
	testutils.CheckString("[5,10,]", dll.String(), t)
	testutils.CheckInt(2, dll.Size(), t)

	val, err = dll.PopBack()
	if err != nil {
		t.Errorf("Got error %s", err.Error())
	} else {
		testutils.CheckInt(10, val, t)
	}
	testutils.CheckString("[5,]", dll.String(), t)
	testutils.CheckInt(1, dll.Size(), t)

	val, err = dll.PopBack()
	if err != nil {
		t.Errorf("Got error %s", err.Error())
	} else {
		testutils.CheckInt(5, val, t)
	}
	testutils.CheckString("[]", dll.String(), t)
	testutils.CheckInt(0, dll.Size(), t)

	val, err = dll.PopBack()
	if err == nil {
		t.Error("Was testutils.Expecting error, got nil")
	}
}

func TestPopFront(t *testing.T) {
	dll := new(DoublyLinkedIntegerList)
	dll.PushBack(5)
	dll.PushBack(10)
	dll.PushBack(36)
	dll.PushBack(42)
	testutils.CheckString("[5,10,36,42,]", dll.String(), t)
	testutils.CheckInt(4, dll.Size(), t)

	val, err := dll.PopFront()
	if err != nil {
		t.Errorf("Got error %s", err.Error())
	} else {
		testutils.CheckInt(5, val, t)
	}
	testutils.CheckString("[10,36,42,]", dll.String(), t)
	testutils.CheckInt(3, dll.Size(), t)

	val, err = dll.PopFront()
	if err != nil {
		t.Errorf("Got error %s", err.Error())
	} else {
		testutils.CheckInt(10, val, t)
	}
	testutils.CheckString("[36,42,]", dll.String(), t)
	testutils.CheckInt(2, dll.Size(), t)

	val, err = dll.PopFront()
	if err != nil {
		t.Errorf("Got error %s", err.Error())
	} else {
		testutils.CheckInt(36, val, t)
	}
	testutils.CheckString("[42,]", dll.String(), t)
	testutils.CheckInt(1, dll.Size(), t)

	val, err = dll.PopFront()
	if err != nil {
		t.Errorf("Got error %s", err.Error())
	} else {
		testutils.CheckInt(42, val, t)
	}
	testutils.CheckString("[]", dll.String(), t)
	testutils.CheckInt(0, dll.Size(), t)

	val, err = dll.PopFront()
	if err == nil {
		t.Error("Was testutils.Expecting error, got nil")
	}
}
