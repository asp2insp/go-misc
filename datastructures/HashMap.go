package datastructures

import "fmt"

const lowerLF = 0.25
const upperLF = 0.75

type entry struct {
	key, value string
}

type HashMap interface {
	Put(string, string)            // O(1)
	Get(string) (string, error)    // O(1)
	Remove(string) (string, error) // O(1)
	Contains(string) bool          // O(1)
	Size() int                     // O(1)
	Capacity() int                 // O(1)
}

type QuadraticProbedHashMap struct {
	data         []entry
	hashFunction func(string) uint
	size         int
}

func NewHashMap(initialSize int, hashFunction func(string) uint) *QuadraticProbedHashMap {
	if initialSize <= 0 {
		initialSize = 4
	}
	ret := new(QuadraticProbedHashMap)
	ret.data = make([]entry, initialSize)
	ret.hashFunction = hashFunction
	return ret
}

func (m *QuadraticProbedHashMap) get(i uint) *entry {
	return &m.data[i%uint(len(m.data))]
}

func (m *QuadraticProbedHashMap) Size() int {
	return m.size
}

func (m *QuadraticProbedHashMap) Capacity() int {
	return len(m.data)
}

func (m *QuadraticProbedHashMap) findSlot(key string) *entry {
	base := m.hashFunction(key)
	e := m.get(base)
	if e.key == key {
		return e
	}
	var i uint
	for i = 1; ; i++ {
		e = m.get(base + i ^ 2)
		if e.key == "" || e.key == key {
			return e
		}
	}
	return nil
}

func (m *QuadraticProbedHashMap) Contains(key string) bool {
	e := m.findSlot(key)
	return e.key == key
}

func (m *QuadraticProbedHashMap) Get(key string) (string, error) {
	e := m.findSlot(key)
	if e.key == "" {
		return "", fmt.Errorf("Key '%s' not found", key)
	} else if e.key == key {
		return e.value, nil
	} else {
		return "", fmt.Errorf("You suck!")
	}
}

func (m *QuadraticProbedHashMap) Put(key, value string) {
	m.put(key, value)
	m.checkSize()
}

func (m *QuadraticProbedHashMap) put(key, value string) {
	e := m.findSlot(key)
	if e.key != key {
		m.size += 1
	}
	e.key = key
	e.value = value
}

func (m *QuadraticProbedHashMap) checkSize() {
	currentLF := float64(m.size) / float64(len(m.data))
	var dataCopy []entry
	if currentLF > upperLF {
		dataCopy = m.data
		m.size = 0
		m.data = make([]entry, len(m.data)*2)
	} else if currentLF < lowerLF {
		dataCopy = m.data
		m.size = 0
		m.data = make([]entry, len(m.data)/2)
	} else {
		return
	}
	for _, e := range dataCopy {
		if e.key != "" {
			m.put(e.key, e.value)
		}
	}
}

func (m *QuadraticProbedHashMap) Remove(key string) (string, error) {
	e := m.findSlot(key)
	if e.key == "" {
		return "", fmt.Errorf("Key '%s' not found", key)
	} else if e.key == key {
		e.key = ""
		m.size -= 1
		return e.value, nil
	} else {
		return "", fmt.Errorf("You suck!")
	}
}
