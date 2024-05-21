package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~|\*|=|-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	total := 0
	for _, line := range lines {
		if re.MatchString(line) {
			total++
		}
	}
	return total
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[\d]+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) (result []string) {
	re := regexp.MustCompile(`User +([a-zA-z0-9]+) `)
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) > 0 {
			line = "[USR] " + match[1] + " " + line
		}
		result = append(result, line)
	}
	return
}
