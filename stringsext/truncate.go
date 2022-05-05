package stringsext

import "strings"

// TruncateText truncates the string to at most the specified length.  It will
// round up to the nearest space in an attempt to avoid truncating words.
func TruncateText(s string, max int) string {
	if max <= 0 {
		return ""
	}
	s = strings.Join(strings.Fields(s), " ")
	if max >= len(s) {
		return s
	}
	li := strings.LastIndex(s[:max], " ")
	if li == -1 {
		return s[:max]
	}
	return s[:li]
}
