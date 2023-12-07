package main

import (
	"testing"
)

func TestIsDigit(t *testing.T) {
	if IsDigit('0') != true {
		t.Errorf("Result was incorrect:\n     got <false>\nexpected <true>")
	}
	if IsDigit('9') != true {
		t.Errorf("Result was incorrect:\n     got <false>\nexpected <true>")
	}
	if IsDigit('a') != false {
		t.Errorf("Result was incorrect:\n     got <true>\nexpected <false>")
	}
	if IsDigit('z') != false {
		t.Errorf("Result was incorrect:\n     got <true>\nexpected <false>")
	}
	if IsDigit('A') != false {
		t.Errorf("Result was incorrect:\n     got <true>\nexpected <false>")
	}
	if IsDigit('Z') != false {
		t.Errorf("Result was incorrect:\n     got <true>\nexpected <false>")
	}
}

func TestGetStrength(t *testing.T) {
	expected := 2
	input := byte('2')
	result := GetStrength(input)
	if result != expected {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}

	expected = 11
	input = byte('J')
	result = GetStrength(input)
	if result != expected {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}
}

func TestGetStrengthJ(t *testing.T) {
	expected := 2
	input := byte('2')
	result := GetStrengthJ(input)
	if result != expected {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}

	expected = 0
	input = byte('J')
	result = GetStrengthJ(input)
	if result != expected {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}
}

func TestReplaceJoker(t *testing.T) {
	expected := "AAA55"
	input := "AJA55"
	result := ReplaceJoker(input)
	if result != expected {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}

	expected = "JJJJJ"
	input = "JJJJJ"
	result = ReplaceJoker(input)
	if result != expected {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}

	expected = "AAAAA"
	input = "JJJJA"
	result = ReplaceJoker(input)
	if result != expected {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}

	expected = "22223"
	input = "22JJ3"
	result = ReplaceJoker(input)
	if result != expected {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}
}

func TestJokerExists(t *testing.T) {
	input := "22J33"
	if JokerExists(&input) != true {
		t.Errorf("Result was incorrect:\n     got <false>\nexpected <true>")
	}

	input = "2234J"
	if JokerExists(&input) != true {
		t.Errorf("Result was incorrect:\n     got <false>\nexpected <true>")
	}

	input = "22345"
	if JokerExists(&input) != false {
		t.Errorf("Result was incorrect:\n     got <true>\nexpected <false>")
	}

	input = "TQKAA"
	if JokerExists(&input) != false {
		t.Errorf("Result was incorrect:\n     got <true>\nexpected <false>")
	}
}
