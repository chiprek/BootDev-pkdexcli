package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " i want to be the very best ",
			expected: []string{"i", "want", "to", "be", "the", "very", "best"},
		},
		{
			input:    "What Do You Mean Mr.Krabs",
			expected: []string{"what", "do", "you", "mean", "mr.krabs"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("length miss match. expected: %d, got: %d", len(c.expected), len(actual))
			t.Fail()
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected: %v, got: %v", expectedWord, word)
				t.Fail()
			}
		}
	}
}
