// Package contains string utilities
package strutil

// #include "strutil.h"
import "C"

// Func check if char occurs once in string
// Estimate time: O(n) Estimate required memory: O(1)
func IsUniqueChars(str string) bool {
	if len(str) > 256 {
		return false
	}

	var charSet [256] bool
	for i := 0; i < len(str); i++ {
		var val int
		val = int(str[i])
		if charSet[val] {
			return false
		}
		charSet[val] = true
	}

	return true
}

// Func that reverse string, which ending null
func Reverse(str string) string {
	ch := C.CString(str)
	C.reverse(ch);
	res := C.GoString(ch);
	return res
}