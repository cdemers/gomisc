package stringsext

import "testing"

func TestTruncateText(t *testing.T) {
	tests := []struct {
		input    string
		length   int
		expected string
	}{
		{"aaaa aaaa", 3, ""},
		{"aaaa aaaa", 4, "aaaa"},
		{"aaaa aaaa", 5, "aaaa"},
		{"aaaa aaaa", 6, "aaaa"},
		{"aaaa aaaa", 9, "aaaa aaaa"},
		{"aaaa aaaa", 10, "aaaa aaaa"},
	}
	for _, test := range tests {
		actual := TruncateText(test.input, test.length)
		if actual != test.expected {
			t.Errorf("TruncateText(%q, %d, %q) = %q; expected %q", test.input, test.length, test.ellipsis, actual, test.expected)
		}
	}
}
