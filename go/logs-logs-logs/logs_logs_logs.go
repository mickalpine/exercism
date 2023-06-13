package logs

import "unicode/utf8"

var applicationMap = map[rune]string{
	'❗': "recommendation",
	'🔍': "search",
	'☀': "weather",
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, rune := range log {
		if app, exists := applicationMap[rune]; exists {
			return app
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runeString := []rune(log)
	for index, rune := range runeString {
		if rune == oldRune {
			runeString[index] = newRune
		}
	}

	return string(runeString)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
