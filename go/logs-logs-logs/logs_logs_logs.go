package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {

	for _, l := range log {
		switch l {
		case '\u2757':
			return "recommendation"
		case '\U0001F50D':
			return "search"
		case '\u2600':
			return "weather"
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {

	// note for later:
	// runes aren't being replaced.
	// maybe try printing usign %c instead
	var newLog []rune
	for _, v := range log {
		if v == oldRune {
			newLog = append(newLog, newRune)
		} else {
			newLog = append(newLog, v)
		}
	}
	return string(newLog)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	c := utf8.RuneCountInString(log)
	if c > limit {
		return false
	}
	return true
}
