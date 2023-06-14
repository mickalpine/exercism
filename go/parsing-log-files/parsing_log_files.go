package parsinglogfiles

import (
	"fmt"
	"regexp"
)

var (
	validLineRegex       *regexp.Regexp = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	splitLogLineRegex    *regexp.Regexp = regexp.MustCompile(`<[-=~*]*>`)
	countQuotedRegex     *regexp.Regexp = regexp.MustCompile(`(?i)".*password.*"`)
	removeEndOfLineRegex *regexp.Regexp = regexp.MustCompile(`end-of-line\d*`)
	tagWithUserNameRegex *regexp.Regexp = regexp.MustCompile(`User\s+(\w+)`)
)

func IsValidLine(text string) bool {
	return validLineRegex.MatchString(text)
}

func SplitLogLine(text string) []string {
	return splitLogLineRegex.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, line := range lines {
		if countQuotedRegex.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	return removeEndOfLineRegex.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	for i, line := range lines {
		match := tagWithUserNameRegex.FindStringSubmatch(line)
		if match != nil {
			name := match[1]
			lines[i] = fmt.Sprintf("[USR] %s %s", name, line)
		}
	}

	return lines
}
