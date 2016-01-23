package testutils

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
	"testing"
)

func s() string {
	buf := make([]byte, 1000)
	n := runtime.Stack(buf, false)
	stack := bytes.Split(buf[:n], []byte("\n"))
	i := 1
	for ; strings.Index(string(stack[i]), "testutils") != -1; i++ {
	}
	return fmt.Sprintf("%s\n%s\n", string(stack[i]), string(stack[i+1]))
}

// CheckString ensures the given strings are equal
func CheckString(exp, obs string, t *testing.T) {
	if exp != obs {
		t.Errorf("%s: Expected %s but got %s", s(), exp, obs)
	}
}

// CheckInt ensures the given ints are equal
func CheckInt(exp, obs int, t *testing.T) {
	if exp != obs {
		t.Errorf("%s: Expected %s but got %s", s(), exp, obs)
	}
}

// CheckInt64 ensures the given int64s are equal
func CheckInt64(exp, obs int64, t *testing.T) {
	if exp != obs {
		t.Errorf("%s: Expected %s but got %s", s(), exp, obs)
	}
}

// CheckUint ensures the given uints are equal
func CheckUint(exp, obs uint, t *testing.T) {
	if exp != obs {
		t.Errorf("%s: Expected %s but got %s", s(), exp, obs)
	}
}

// CheckUint ensures the given uints are equal
func CheckUint32(exp, obs uint32, t *testing.T) {
	if exp != obs {
		t.Errorf("%s: Expected %s but got %s", s(), exp, obs)
	}
}

// CheckUint64 ensures the given uint64s are equal
func CheckUint64(exp, obs uint64, t *testing.T) {
	if exp != obs {
		t.Errorf("%s: Expected %s but got %s", s(), exp, obs)
	}
}

// CheckByte ensures the given bytes are equal
func CheckByte(exp, obs byte, t *testing.T) {
	if exp != obs {
		t.Errorf("%s: Expected %s but got %s", s(), exp, obs)
	}
}

// ExpectTrue ensures the given resolved expression is true
func ExpectTrue(obs bool, message string, t *testing.T) {
	if !obs {
		t.Errorf("%s: Expected true but got false: %s", s(), message)
	}
}

// ExpectFalse ensures the given resolved expresion is false
func ExpectFalse(obs bool, message string, t *testing.T) {
	if obs {
		t.Errorf("%s: Expected false but got true: %s", s(), message)
	}
}

// ExpectNil ensures the given expression is nil
func ExpectNil(obs interface{}, message string, t *testing.T) {
	if obs != nil {
		t.Errorf("%s: Expected nil but got: %v. (%s)", s(), obs, message)
	}
}

// CheckErr converts a runtime error into a testing error
func CheckErr(err error, t *testing.T) {
	if err != nil {
		t.Error(err.Error())
		t.Error(s())
		t.FailNow()
	}
}

// CheckSlice ensures that the two int slices are equal
func CheckSlice(exp, obs []int, t *testing.T) {
	if len(exp) != len(obs) {
		t.Errorf("%s: Differing slice lengths: %d vs %d", s(), len(exp), len(obs))
	}
	for i, v := range exp {
		if v != obs[i] {
			t.Errorf("%s: Slices differ in element number %d: %d vs %d", s(), i, v, obs[i])
		}
	}
}

// CheckByteSlice ensures that the two byte slices are equal
func CheckByteSlice(exp, obs []byte, t *testing.T) {
	if len(exp) != len(obs) {
		t.Errorf("%s: Differing slice lengths: %d vs %d", s(), len(exp), len(obs))
	}
	for i, v := range exp {
		if v != obs[i] {
			t.Errorf("%s: Slices differ in element number %d: %d vs %d", s(), i, v, obs[i])
		}
	}
}

// CheckStringSlice ensures that the two string slices are equal
func CheckStringSlice(exp, obs []string, t *testing.T) {
	if len(exp) != len(obs) {
		t.Errorf("%s: Differing slice lengths: %d vs %d", s(), len(exp), len(obs))
	}
	for i, v := range exp {
		if v != obs[i] {
			t.Errorf("%s: Slices differ in element number %d: %v vs %v", s(), i, v, obs[i])
		}
	}
}
