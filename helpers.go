package irc

import "strings"

var (
	tagEscape   = strings.NewReplacer(";", `\:`, " ", `\s`, "\r", `\r`, "\n", `\n`, "\\", `\\`)
	tagUnescape = strings.NewReplacer(`\:`, ";", `\s`, " ", `\r`, "\r", `\n`, "\n", `\\`, "\\")
)

// Hand written string.Fields, much faster than string.FieldsFunc.
func stringFields(s string, sep byte) []string {
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

	return out
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
