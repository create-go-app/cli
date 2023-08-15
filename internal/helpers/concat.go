package helpers

import (
	"unsafe"
)

// Concat concatenate strings using the built-in copy and "unsafe" package with
// unsafe.String function.
func Concat(s ...string) string {
	if len(s) == 0 {
		return ""
	}

	n := 0
	for i := 0; i < len(s); i++ {
		n += len(s[i])
	}

	b := make([]byte, n)

	idx := 0
	for i := 0; i < len(s); i++ {
		copy(b[idx:], s[i])
		idx += len(s[i])
	}

	return unsafe.String(unsafe.SliceData(b), n)
}
