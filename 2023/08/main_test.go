package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	expected_k := "AAA"
	expected_v := []string{"BBB", "CCC"}

	input := "AAA = (BBB, CCC)"
	result_k, result_v := ParseInput(&input)

	if result_k != expected_k {
		t.Errorf("Result was incorrect:\ngot     <%v>\nexpected<%v>",
			result_k,
			expected_k)
	}

	if !reflect.DeepEqual(result_v, expected_v) {
		t.Errorf("Result was incorrect:\ngot     <%v>\nexpected<%v>",
			result_v,
			expected_v)
	}
}

func TestGcd(t *testing.T) {
	expected := 5
	result := Gcd(10, 5)
	if result != expected {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}
}

func TestLcm(t *testing.T) {
	expected := 30
	result := Lcm(10, 15)
	if result != expected {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}

	expected = 60
	result = Lcm(10, 15, 20)
	if result != expected {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}

	expected = 2520
	result = Lcm(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	if result != expected {
		t.Errorf("Result was incorrect:\n     got <%v>\nexpected <%v>",
			result,
			expected)
	}
}
