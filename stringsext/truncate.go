package stringsext

import "strings"

func TruncateText(s string, max int) string {
	s = strings.Join(strings.Fields(s), " ")
	if max > len(s) {
		return s
	}
	return s[:strings.LastIndex(s[:max], " ")]
}
