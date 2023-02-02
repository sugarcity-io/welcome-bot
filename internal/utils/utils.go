package utils

import (
	"strings"
)

// Check if a string starts with a prefix.
func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}
