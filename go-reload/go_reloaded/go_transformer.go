package go_reloaded

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// UPPERCASE
func ToUp(str string) string {
	var b strings.Builder
	for _, c := range str {
		b.WriteRune(unicode.ToUpper(c))
	}
	return b.String()
}

// lowercase
func ToLow(str string) string {
	var b strings.Builder
	for _, c := range str {
		b.WriteRune(unicode.ToLower(c))
	}
	return b.String()
}

// Capitalize: first letter uppercase, rest lowercase
func ToCap(str string) string {
	var b strings.Builder
	first := true

	for _, c := range str {
		if first {
			b.WriteRune(unicode.ToUpper(c))
			first = false
		} else {
			b.WriteRune(unicode.ToLower(c))
		}
	}
	return b.String()
}

// Convert a decimal number string → binary string
func ToBin(str string) string {
	n, err := strconv.Atoi(str)
	if err != nil {
		return "" // or return error if you prefer
	}
	return strconv.FormatInt(int64(n), 2)
}

// Convert a decimal number string → hexadecimal string
func ToHex(str string) string {
	n, err := strconv.Atoi(str)
	if err != nil {
		return ""
	}
	fmt.Println(n)
	return strconv.FormatInt(int64(n), 16)
}
