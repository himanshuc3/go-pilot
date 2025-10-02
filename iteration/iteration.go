package iteration

import "strings"

func Repeat(character string) string {
	// NOTE
	// 1. `strings.Builder` is a more efficient way to concatenate strings in Go
	// because of memory allocation.
	var repeated strings.Builder
	for i := 0; i < 5; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
