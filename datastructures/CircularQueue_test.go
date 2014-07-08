package datastructures

import "testing"
import "github.com/ca-geo/interview-prep/testutils"

func TestSatisfiesInterface(t *testing.T) {
	var cq CircularQueue
	acq := new(ArrayCircularQueue)
	cq = acq
	testutils.CheckInt(0, cq.Size(), t)
}

func TestBasics(t *testing.T) {
	acq := CreateQueue()
	acq.Enqueue("John")
	acq.Enqueue("Jacob")
	acq.Enqueue("Jingleheimer")
	acq.Enqueue("Schmidt")
	acq.Enqueue("his")
	acq.Enqueue("name")
	acq.Enqueue("is")

	testutils.CheckInt(7, acq.Size(), t)

	s, err := acq.Dequeue()
	testutils.CheckErr(err, t)
	testutils.CheckString("John", s.(string), t)
	s, err = acq.Dequeue()
	testutils.CheckErr(err, t)
	testutils.CheckString("Jacob", s.(string), t)
	s, err = acq.Dequeue()
	testutils.CheckErr(err, t)
	testutils.CheckString("Jingleheimer", s.(string), t)
	s, err = acq.Dequeue()
	testutils.CheckErr(err, t)
	testutils.CheckString("Schmidt", s.(string), t)

	testutils.CheckInt(3, acq.Size(), t)

	acq.Enqueue("my")
	acq.Enqueue("name")
	acq.Enqueue("too")

	testutils.CheckInt(6, acq.Size(), t)
	testutils.CheckString("[his,name,is,my,name,too]", acq.String(), t)
}
