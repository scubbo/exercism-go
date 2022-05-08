package isogram

import "unicode"

func IsIsogram(word string) bool {
	// Really we would want this to be a Set, but no such type exists
	// in GoLang's standard libraries AFAICT
	present_letters := make(map[rune]bool)
	for _, l := range word {
		// No built-in `contains`!? Wow.
		// https://gosamples.dev/slice-contains/
		if l == '-' || l == ' ' {
			continue
		}
		if present_letters[unicode.ToLower(l)] {
			return false
		}
		present_letters[unicode.ToLower(l)] = true
	}
	return true
}
