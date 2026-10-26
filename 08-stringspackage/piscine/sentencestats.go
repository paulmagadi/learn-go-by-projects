package piscine

import (
	"strings"
)	

func SentenceStats(s string) (words int, chars int, hasHello bool) {
	fields := strings.Fields(s) // Split the string into fields (words) using whitespace as the delimiter
	lenfields := len(fields) // Count the number of fields (words)
	stringlen := len(s) // Count the number of characters in the string
	stringToLower := strings.ToLower(s) // Convert the string to lowercase for case-insensitive comparison
	hasHello = strings.Contains(stringToLower, "hello") // Check if the lowercase string contains "hello"

	return lenfields, stringlen, hasHello
}