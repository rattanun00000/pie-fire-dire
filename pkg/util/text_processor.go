package util

import (
	"regexp"
	"strings"
)

func CleanText(text string) string {
	text = strings.ToLower(text)

	text = regexp.MustCompile(`[.,;:!?]`).ReplaceAllString(text, " ")

	return text
}
