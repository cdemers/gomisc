package stringsext

import "strings"

func TruncateText(s string, max int) string {
	if max <= 0 {
		return ""
	}
	s = strings.Join(strings.Fields(s), " ")
	if max > len(s) {
		return s
	}
	li := strings.LastIndex(s[:max], " ")
	if li == -1 {
		return s[:max]
	}
	return s[:li]
}
