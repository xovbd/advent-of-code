package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	expected := []int{-11, -15, 19, 0, 5, -31}
	input := "-11 -15 19 0 5 -31"
	result := ParseInput(&input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result was incorrect:\ngot     <%v>\nexpected<%v>",
			result,
			expected)
	}
}

func TestExtrapolateSeq(t *testing.T) {
	expected := -625
	input := []int{-11, -15, 19, 0, 5, -31}
	result := ExtrapolateSeq(input)

	if result != expected {
		t.Errorf("Result was incorrect:\ngot     <%v>\nexpected<%v>",
			result,
			expected)
	}

	expected = 18
	input = []int{0, 3, 6, 9, 12, 15}
	result = ExtrapolateSeq(input)

	if result != expected {
		t.Errorf("Result was incorrect:\ngot     <%v>\nexpected<%v>",
			result,
			expected)
	}

	expected = 28
	input = []int{1, 3, 6, 10, 15, 21}
	result = ExtrapolateSeq(input)

	if result != expected {
		t.Errorf("Result was incorrect:\ngot     <%v>\nexpected<%v>",
			result,
			expected)
	}

	expected = 68
	input = []int{10, 13, 16, 21, 30, 45}
	result = ExtrapolateSeq(input)

	if result != expected {
		t.Errorf("Result was incorrect:\ngot     <%v>\nexpected<%v>",
			result,
			expected)
	}
}
