package luhn

import (
	"regexp"
	"strconv"
)

var stringStripRegex = regexp.MustCompile(" ")
var anyNonDigitRegex = regexp.MustCompile(`\D`)

func Valid(id string) bool {

	id = stringStripRegex.ReplaceAllString(id, "")
	if len(id) < 2 {
		return false
	}
	if anyNonDigitRegex.MatchString(id) {
		return false
	}

	total := 0
	for i := 0; i < len(id); i++ {
		r := id[len(id)-(i+1)]
		if r == ' ' {
			continue
		}
		val, err := strconv.Atoi(string(r))
		if err != nil {
			return false
		}
		if i%2 == 1 {
			val = doubleAndKeepUnderTen(val)
		}
		total += val
	}
	return total%10 == 0

}

func doubleAndKeepUnderTen(num int) int {
	doubled := num * 2
	if doubled > 9 {
		doubled -= 9
	}
	return doubled
}
