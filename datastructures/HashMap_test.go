package datastructures

import "testing"
import "github.com/ca-geo/interview-prep/testutils"

func hash(s string) uint {
	var hash uint = 5381

	for i := 0; i < len(s); i++ {
		hash = ((hash << 5) + hash) + uint(s[i])
	}
	return hash
}

func TestPut(t *testing.T) {
	m := NewHashMap(2, hash)
	testutils.CheckInt(0, m.Size(), t)

	m.Put("John", "John Doe")
	testutils.CheckInt(1, m.Size(), t)

	m.Put("Fred", "Fred Wilson")
	testutils.CheckInt(2, m.Size(), t)

	testutils.ExpectTrue(m.Contains("John"), "Map should contain John", t)
	testutils.ExpectTrue(m.Contains("Fred"), "Map should contain Fred", t)
	testutils.ExpectFalse(m.Contains("George"), "Map shouldn't contain George", t)
}

func TestMapRemove(t *testing.T) {
	m := NewHashMap(3, hash)

	m.Put("Red", "Blood")
	m.Put("Green", "Money")
	m.Put("Black", "Night")
	m.Put("Yellow", "Sun")

	testutils.CheckInt(4, m.Size(), t)
	testutils.ExpectTrue(m.Contains("Red"), "Map should contain Red", t)

	m.Remove("Red")
	testutils.ExpectFalse(m.Contains("Red"), "Red was removed from map", t)
	testutils.CheckInt(3, m.Size(), t)

	testutils.ExpectTrue(m.Contains("Yellow"), "Map should contain Yellow", t)
	m.Remove("Yellow")
	testutils.ExpectFalse(m.Contains("Yellow"), "Yellow was removed from map", t)
	testutils.CheckInt(2, m.Size(), t)

	testutils.ExpectTrue(m.Contains("Black"), "Map should contain Black", t)
	m.Remove("Black")
	testutils.ExpectFalse(m.Contains("Black"), "Black was removed from map", t)
	testutils.CheckInt(1, m.Size(), t)

	testutils.ExpectFalse(m.Contains("Orange"), "Map never had Orange", t)
	_, err := m.Remove("Orange")
	testutils.ExpectTrue(err != nil, "Should get an error for key Orange", t)

	testutils.ExpectTrue(m.Contains("Green"), "Map should contain Green", t)
	m.Remove("Green")
	testutils.ExpectFalse(m.Contains("Green"), "Green was removed from map", t)
	testutils.CheckInt(0, m.Size(), t)
}
