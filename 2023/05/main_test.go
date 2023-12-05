package main

import (
	"reflect"
	"testing"
)

func TestExtractSeeds(t *testing.T) {
	expected := []uint64{79, 14, 55, 13}
	input := "seeds: 79 14 55 13"
	result := ExtractSeeds(&input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}
}

func TestExtractTopic(t *testing.T) {
	expected := "seed-to-soil"
	input := "seed-to-soil map:"
	result := ExtractTopic(&input)
	if result != expected {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}
}

func TestExtractRule(t *testing.T) {
	expected := Rule{50, 98, 2}
	input := "50 98 2"
	result := ExtractRule(&input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}
}

func TestGetValueSrc(t *testing.T) {
	rules := []Rule{{50, 98, 2}, {52, 50, 48}}

	expected := uint64(81)
	input := uint64(79)
	result := GetValueSrc(input, rules)
	if result != expected {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}

	expected = uint64(14)
	input = uint64(14)
	result = GetValueSrc(input, rules)
	if result != expected {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}

	expected = uint64(57)
	input = uint64(55)
	result = GetValueSrc(input, rules)
	if result != expected {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}

	expected = uint64(13)
	input = uint64(13)
	result = GetValueSrc(input, rules)
	if result != expected {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}
}

func TestGetValueDst(t *testing.T) {
	rules := []Rule{{50, 98, 2}, {52, 50, 48}}

	expected := uint64(79)
	input := uint64(81)
	result := GetValueDst(input, rules)
	if result != expected {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}

	expected = uint64(14)
	input = uint64(14)
	result = GetValueDst(input, rules)
	if result != expected {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}

	expected = uint64(99)
	input = uint64(51)
	result = GetValueDst(input, rules)
	if result != expected {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}

	expected = uint64(13)
	input = uint64(13)
	result = GetValueDst(input, rules)
	if result != expected {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}
}
