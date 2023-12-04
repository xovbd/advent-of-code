package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	expected := Card{
		id:        1,
		numsWin:   []string{"41", "48", "83", "86", "17"},
		numsCheck: []string{"83", "86", "6", "31", "17", "9", "48", "53"},
		matches:   4,
		points:    8,
		quantity:  1,
	}

	input := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	result := ParseInput(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result was incorrect:\ngot     <%v>\nexpected<%v>",
			result,
			expected)
	}

	input = "Card   1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	result = ParseInput(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result was incorrect:\ngot     <%v>\nexpected<%v>",
			result,
			expected)
	}
}
