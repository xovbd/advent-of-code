package main

import (
	"reflect"
	"testing"
)

func TestExtractValue(t *testing.T) {
	expected := uint64(256133211111425)
	input := "Foo:   256   1332   1111 1425"
	result := ExtractValue(&input)
	if result != expected {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}
}

func TestExtractValues(t *testing.T) {
	expected := []uint64{256, 1332, 1111, 1425}
	input := "Foo:   256   1332   1111 1425"
	result := ExtractValues(&input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}
}

func TestDetermineOpportunities(t *testing.T) {
	expected := []uint64{4, 8, 9}
	result := DetermineOpportunities(&[]uint64{7, 15, 30}, &[]uint64{9, 40, 200})
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}
}

func TestSolvePart1(t *testing.T) {
	expected := uint64(288)
	result := SolvePart1(&[]uint64{4, 8, 9})
	if result != expected {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}
}

func TestSolvePart2(t *testing.T) {
	expected := uint64(71503)
	time := uint64(71530)
	dist := uint64(940200)
	result := SolvePart2(&time, &dist)
	if result != expected {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}
}
