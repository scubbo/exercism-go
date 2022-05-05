package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	validLineRegex := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return validLineRegex.MatchString(text)
}

func SplitLogLine(text string) []string {
	splitRegex := regexp.MustCompile(`<[~*=-]*>`)
	return splitRegex.Split(text, -1)
}

func CountQuotedPasswords(lines []string) (count int) {
	// "Each line will contain at most two quotation marks", so we don't
	// need to worry about cases like
	//`"this is quote-enclosed" this is not, but contains  password "more quote-enclosed text"`
	matchRegex := regexp.MustCompile(`(?i)".*password.*"`)
	for _, s := range lines {
		if matchRegex.MatchString(s) {
			count++
		}
	}
	return
}

func RemoveEndOfLineText(text string) string {
	endOfLineRegex := regexp.MustCompile(`end-of-line\d*`)
	return endOfLineRegex.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) (output []string) {
	matchRegex := regexp.MustCompile(`User\s+(\S*)`)
	for _, line := range lines {
		match := matchRegex.FindStringSubmatch(line)
		if match != nil {
			output = append(output, "[USR] "+match[1]+" "+line)
		} else {
			output = append(output, line)
		}
	}
	return
}
