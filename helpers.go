package irc

import (
	"io"
	"strings"
)

// Hand written string.Fields, much faster than string.FieldsFunc.
// Note that unlike strings.Fields, if the returned slice would be empty, it
// will be nil.
func stringFields(s string, sep byte) []string {
	s = fastTrim(s, sep)

	if s == "" {
		return nil
	}

	count := strings.Count(s, string(sep))
	if count == 0 {
		return []string{s}
	}

	out := make([]string, 0, count+1)

	for {
		i := strings.IndexByte(s, sep)
		if i == -1 {
			break
		}

		f := s[:i]
		s = s[i+1:]

		if f != "" {
			out = append(out, f)
		}
	}

	if s != "" {
		out = append(out, s)
	}

	if len(out) == 0 {
		return nil
	}

	return out
}

// fastTrim is a faster implementation of trimming a string against a single
// byte.
func fastTrim(s string, b byte) string {
	if s == "" {
		return ""
	}

	br := rune(b)

	findLeft := true
	left := 0

	prevB := false
	right := 0

	for i, r := range s {
		if findLeft {
			if r != br {
				left = i
				right = i
				findLeft = false
			}
			continue
		}

		if r == br {
			if prevB {
				continue
			}
			prevB = true
			right = i
		} else {
			prevB = false
			right = i
		}
	}

	if !prevB {
		right++
	}

	return s[left:right]
}

var (
	tagEscapeReplacer   = strings.NewReplacer(";", `\:`, " ", `\s`, "\r", `\r`, "\n", `\n`, "\\", `\\`)
	tagUnescapeReplacer = strings.NewReplacer(`\:`, ";", `\s`, " ", `\r`, "\r", `\n`, "\n", `\\`, "\\")
)

func tagUnescape(s string) string {
	// Check ahead of time to see if the string contains something that
	// needs to be unescaped. This is faster and much more allocation
	// efficient than always running the replacer.
	if containsUnescapeable(s) {
		return tagUnescapeReplacer.Replace(s)
	}
	return s
}

type stringWriter interface {
	io.Writer
	WriteString(string) (int, error)
}

func tagEscapeWrite(w stringWriter, s string) (n int, err error) {
	if containsEscapable(s) {
		return tagEscapeReplacer.WriteString(w, s)
	}
	return w.WriteString(s)
}

func containsEscapable(s string) bool {
	for _, r := range s {
		switch r {
		case ';', ' ', '\n', '\r', '\\':
			return true
		}
	}
	return false
}

func countEscapeable(s string) int {
	count := 0

	for _, r := range s {
		switch r {
		case ';', ' ', '\n', '\r', '\\':
			count++
		}
	}

	return count
}

func containsUnescapeable(s string) bool {
	bi := strings.IndexByte(s, '\\') + 1
	if bi != 0 && bi < len(s) {
		switch s[bi] {
		case ':', 's', 'n', 'r', '\\':
			return true
		}
	}

	return false
}
