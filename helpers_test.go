package irc

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringFieldsHelper(t *testing.T) {
	tests := []struct {
		s        string
		sep      byte
		expected []string
	}{
		{
			s:   "",
			sep: ' ',
		},
		{
			s:        "A",
			sep:      ' ',
			expected: []string{"A"},
		},
		{
			s:        "A B C",
			sep:      ' ',
			expected: []string{"A", "B", "C"},
		},
		{
			s:        "A B C",
			sep:      '-',
			expected: []string{"A B C"},
		},
		{
			s:        "A-B-C",
			sep:      '-',
			expected: []string{"A", "B", "C"},
		},
		{
			s:        "A  B   C    ",
			sep:      ' ',
			expected: []string{"A", "B", "C"},
		},
		{
			s:   "       ",
			sep: ' ',
		},
		{
			s:        "A ",
			sep:      ' ',
			expected: []string{"A"},
		},
	}

	for _, test := range tests {
		builtin := strings.FieldsFunc(test.s, func(r rune) bool {
			return r == rune(test.sep)
		})
		if len(builtin) == 0 {
			builtin = nil
		}
		assert.Equal(t, test.expected, builtin, "expected should match string.Fields")

		got := stringFields(test.s, test.sep)
		assert.Equal(t, test.expected, got, "expected result with sep `%s`", string(test.sep))
	}
}

func TestFastTrimHelper(t *testing.T) {
	tests := []string{
		"",
		"A",
		" A",
		"  A",
		" A ",
		"  A ",
		"  A  ",
		" A  ",
		"A ",
		"A   ",
		"foo bar",
		"   foo bar ",
		"ab ",
		" ab ",
		"ab",
		" ab",
		"a-b-c ",
	}

	for _, test := range tests {
		expected := strings.Trim(test, " ")
		actual := fastTrim(test, ' ')

		assert.Equal(t, expected, actual, "test=`%s`", test)
	}
}
