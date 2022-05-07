package scrabble

import "unicode"

var LetterScore = map[rune]int{
	'D': 2,
	'G': 2,
	'B': 3,
	'C': 3,
	'M': 3,
	'P': 3,
	'F': 4,
	'H': 4,
	'V': 4,
	'W': 4,
	'Y': 4,
	'K': 5,
	'J': 8,
	'X': 8,
	'Q': 10,
	'Z': 10,
}

func Score(word string) (totalScore int) {
	for _, c := range word {
		letterScore := LetterScore[unicode.ToUpper(c)]
		if letterScore == 0 {
			// 0 is the default value ("rune not found"),
			// and 1 is default letter score.
			letterScore = 1
		}
		totalScore += letterScore
	}
	return
}
