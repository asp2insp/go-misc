package testutils

import "testing"

// CheckString ensures the given strings are equal
func CheckString(exp, obs string, t *testing.T) {
	if exp != obs {
		t.Errorf("Expected %s but got %s", exp, obs)
	}
}

// CheckInt ensures the given ints are equal
func CheckInt(exp, obs int, t *testing.T) {
	if exp != obs {
		t.Errorf("Expected %v but got %v", exp, obs)
	}
}

// CheckInt64 ensures the given int64s are equal
func CheckInt64(exp, obs int64, t *testing.T) {
	if exp != obs {
		t.Errorf("Expected %v but got %v", exp, obs)
	}
}

// CheckUint ensures the given uints are equal
func CheckUint(exp, obs uint, t *testing.T) {
	if exp != obs {
		t.Errorf("Expected %v but got %v", exp, obs)
	}
}

// CheckUint64 ensures the given uint64s are equal
func CheckUint64(exp, obs uint64, t *testing.T) {
	if exp != obs {
		t.Errorf("Expected %v but got %v", exp, obs)
	}
}

// ExpectTrue ensures the given resolved expression is true
func ExpectTrue(obs bool, message string, t *testing.T) {
	if !obs {
		t.Errorf("Expected true but got false: %s", message)
	}
}

// ExpectFalse ensures the given resolved expresion is false
func ExpectFalse(obs bool, message string, t *testing.T) {
	if obs {
		t.Errorf("Expected false but got true: %s", message)
	}
}

// ExpectNil ensures the given expression is nil
func ExpectNil(obs interface{}, message string, t *testing.T) {
	if obs != nil {
		t.Errorf("Expected nil but got: %v. (%s)", obs, message)
	}
}

// CheckErr converts a runtime error into a testing error
func CheckErr(err error, t *testing.T) {
	if err != nil {
		t.Error(err.Error())
	}
}

// CheckSlice ensures that the two int slices are equal
func CheckSlice(exp, obs []int, t *testing.T) {
	if len(exp) != len(obs) {
		t.Errorf("Differing slice lengths: %d vs %d", len(exp), len(obs))
	}
	for i, v := range exp {
		if v != obs[i] {
			t.Errorf("Slices differ in element number %d: %d vs %d", i, v, obs[i])
		}
	}
	t.Logf("Expected: %v\nActual: %v", exp, obs)
}

// CheckByteSlice ensures that the two byte slices are equal
func CheckByteSlice(exp, obs []byte, t *testing.T) {
	if len(exp) != len(obs) {
		t.Errorf("Differing slice lengths: %d vs %d", len(exp), len(obs))
	}
	for i, v := range exp {
		if v != obs[i] {
			t.Errorf("Slices differ in element number %d: %d vs %d", i, v, obs[i])
		}
	}
	t.Logf("Expected: %v\nActual: %v", exp, obs)
}

// CheckStringSlice ensures that the two string slices are equal
func CheckStringSlice(exp, obs []string, t *testing.T) {
	if len(exp) != len(obs) {
		t.Errorf("Differing slice lengths: %d vs %d", len(exp), len(obs))
	}
	for i, v := range exp {
		if v != obs[i] {
			t.Errorf("Slices differ in element number %d: %v vs %v", i, v, obs[i])
		}
	}
	t.Logf("Expected: %v\nActual: %v", exp, obs)
}
