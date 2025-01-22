package main

import (
	"strings"
	"unicode"
)

func RemoveNonDigits(s string) string {
	var b strings.Builder
	for _, c := range s {
		if unicode.IsDigit(c) {
			b.WriteRune(c)
		}
	}
	return b.String()
}

func IsRepetitive(s string) bool {
	c := s[0]
	for _, v := range []byte(s) {
		if c != v {
			return false
		}
	}
	return true
}
