package utils

import (
	"testing"
)

func TestHasPrefix(t *testing.T) {
	tests := []struct {
		s      string
		prefix string
		want   bool
	}{
		{"Hello, world!", "Hello", true},
		{"Hello, world!", "Goodbye", false},
		{"", "", true},
		{"Hello, world!", "Hello, world!", true},
	}
	for _, test := range tests {
		if got := HasPrefix(test.s, test.prefix); got != test.want {
			t.Errorf("HasPrefix(%q, %q) = %v, want %v", test.s, test.prefix, got, test.want)
		}
	}
}
