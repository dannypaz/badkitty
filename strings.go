package badkitty

import (
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Contains ..
func Contains(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

// LowerFirstRune is the Go version to .toLowerCase() ;)
func LowerFirstRune(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

// Split ...
func Split(a interface{}, b int) [][]interface{} {
	val := reflect.ValueOf(a)

	origLen := val.Len()
	outerLen := origLen / b
	if origLen%b > 0 {
		outerLen++
	}

	// Make the output slices with all the correct lengths
	c := make([][]interface{}, outerLen)

	for i := range c {
		newLen := b
		if origLen-(i*b+newLen) < 0 {
			newLen = origLen % b
		}
		c[i] = make([]interface{}, newLen)
		for j := range c[i] {
			c[i][j] = val.Index(i*b + j).Interface()
		}
	}

	return c
}

// LeftPad ...
func LeftPad(s string, padStr string, pLen int) string {
	return strings.Repeat(padStr, pLen) + s
}
