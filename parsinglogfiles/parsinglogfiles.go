package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re, _ := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re, _ := regexp.Compile(`<[\~\*\=\-]*>`)
	text = re.ReplaceAllString(text, "-")
	return strings.Split(text, "-")

}

func CountQuotedPasswords(lines []string) int {
	re, _ := regexp.Compile(`(?i)".*password.*"`)
	var count int
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re, _ := regexp.Compile(`end-of-line+\d*`)
	text = re.ReplaceAllString(text, "")
	return text

}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)

	taggedLines := []string{}

	for _, line := range lines {
		if matches := re.FindStringSubmatch(line); matches != nil {
			userName := matches[1]
			taggedLine := fmt.Sprintf("[USR] %s %s", userName, line)
			taggedLines = append(taggedLines, taggedLine)
		} else {
			taggedLines = append(taggedLines, line)
		}
	}

	return taggedLines
}
