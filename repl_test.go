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
			input:    "hellO       mY NAME       is blabla   ",
			expected: []string{"hello", "my", "name", "is", "blabla"},
		},
		{
			input:    "So | sick",
			expected: []string{"so", "|", "sick"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		actualLength := len(actual)
		expectedLength := len(c.expected)
		if actualLength != expectedLength {
			t.Errorf(`
actual length   : %d words
expected length : %d words
`, actualLength, expectedLength)
			return
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf(`
word     : %s 
expected : %s
`, word, expectedWord)
			}
		}
	}
}
