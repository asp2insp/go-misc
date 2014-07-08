package algorithms

import "testing"

func TestReverse(t *testing.T) {
	input := uint32(12345)
	expected := uint32(54321)
	output := ReverseInteger(input)
	if expected != output {
		t.Errorf("Expected %d, got %d", expected, output)
	}

	input = uint32(0)
	expected = uint32(0)
	output = ReverseInteger(input)
	if expected != output {
		t.Errorf("Expected %d, got %d", expected, output)
	}

	input = uint32(3)
	expected = uint32(3)
	output = ReverseInteger(input)
	if expected != output {
		t.Errorf("Expected %d, got %d", expected, output)
	}

	input = uint32(32500)
	expected = uint32(523)
	output = ReverseInteger(input)
	if expected != output {
		t.Errorf("Expected %d, got %d", expected, output)
	}
}
