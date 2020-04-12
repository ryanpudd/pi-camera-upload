package main

import (
	"strings"
)

// Between returns the string within value between a and b
//         returns nothing if a or b doesn't exist
func Between(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)

	posLast := strings.Index(value[posFirstAdjusted:len(value)], b)
	if posLast == -1 {
		return ""
	}
	posLastAdjusted := posFirstAdjusted + posLast

	if posFirstAdjusted >= posLastAdjusted {
		return ""
	}
	return value[posFirstAdjusted:posLastAdjusted]
}

// Before returns the value before a within value or an empty string
func Before(value string, a string) string {
	// Get substring before a string.
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}

// After returns the value after a within value or an empty string
func After(value string, a string) string {
	// Get substring after a string.
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:len(value)]
}
