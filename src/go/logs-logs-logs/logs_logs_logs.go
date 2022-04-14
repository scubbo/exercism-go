package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	lookupTable := map[rune]string{
		'\u2757': "recommendation",
		'🔍':      "search", // Couldn't find the appropriate 4-number expression
		'\u2600': "weather",
	}
	for _, char := range log {
		ident, exists := lookupTable[char]
		if exists {
			return ident
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) (output string) {
	// Unclear from specs whether we should replace in-place, or output afresh.
	// I picked the latter because modifying the structure that you're iterating
	// over generally leads to A Bad Time.
	for _, char := range log {
		if char == oldRune {
			output += string(newRune)
		} else {
			output += string(char)
		}
	}
	return
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
