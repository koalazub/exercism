package parsinglogfiles

import (
	"fmt"
	"regexp"
)

var PrefixedLabels = []string{`^\[TRC\]|^\[DBG\]|^\[INF\]|^\[WRN\]|^\[ERR\]|^\[FTL\]`}

func IsValidLine(text string) bool {
	tokens := `^\[TRC\]|^\[DBG\]|^\[INF\]|^\[WRN\]|^\[ERR\]|^\[FTL\]`
	re := regexp.MustCompile(tokens)
	if re.MatchString(text) {
		return true
	}
	return false
}

// debug function
func compareStrings(srcText, testText []string) {
	minLength := len(srcText)
	if len(testText) < minLength {
		minLength = len(testText)
	}
	for i := 0; i < minLength; i++ {
		f, v := srcText[i], testText[i]
		if len(f) != len(v) {
			fmt.Printf("Length mismatch: fmtText: %d, testText: %d\n", len(f), len(v))
			continue
		}

		if f != v {
			for idx := range v {
				if idx >= len(f) || f[idx] != v[idx] {
					fmt.Printf("Difference at index %d: expected %c, got %c\n", idx, v[idx], f[idx])
					continue
				}
			}
		}
	}

	if len(srcText) > len(testText) {
		fmt.Println("Extra elements in fmtText:", srcText[minLength:])
	} else if len(testText) > len(srcText) {
		fmt.Println("Missing elements in fmtText:", testText[minLength:])
	}

}

// anything with a < and > should immediately  be stripped
func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[-~=\*]*>`)
	splt := re.Split(text, -1)
	return splt
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`"(?i).*password*."`)
	c := 0
	for _, l := range lines {
		if re.MatchString(l) {
			c++
		}
	}
	return c

}

// "[INF] end-of-line23033 Network Failure end-of-line27"
// => "[INF]  Network Failure "
func RemoveEndOfLineText(text string) string {
	if !IsValidLine(text) {
		return ""
	}

	re := regexp.MustCompile(`end-of-line\d+`)
	n := re.ReplaceAllString(text, "")

	return n
}

// "[WRN] User James123 has exceeded storage space.",
//
//	"[USR] James123 [WRN] User James123 has exceeded storage space.",
func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
	var taggedLines []string

	for _, l := range lines {
		match := re.FindStringSubmatch(l)
		if len(match) > 1 {
			taggedLine := fmt.Sprintf("[USR] %s %s", match[1], l)
			taggedLines = append(taggedLines, taggedLine)
		} else {
			taggedLines = append(taggedLines, l)
		}
	}

	return taggedLines
}
