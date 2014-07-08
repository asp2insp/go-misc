package testutils

import "testing"

func CheckString(exp, obs string, t *testing.T) {
	if exp != obs {
		t.Errorf("Expected %s but got %s", exp, obs)
	}
}

func CheckInt(exp, obs int, t *testing.T) {
	if exp != obs {
		t.Errorf("Expected %v but got %v", exp, obs)
	}
}

func ExpectTrue(obs bool, message string, t *testing.T) {
	if !obs {
		t.Errorf("Expected true but got false: %s", message)
	}
}

func ExpectFalse(obs bool, message string, t *testing.T) {
	if obs {
		t.Errorf("Expected false but got true: %s", message)
	}
}

func CheckErr(err error, t *testing.T) {
	if err != nil {
		t.Error(err.Error())
	}
}

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
