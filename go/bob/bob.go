package bob

import (
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	if isQuestion(remark) {
		return "Sure."
	}

	if isYelling(remark) {
		return "Whoa, chill out!"
	}

	if isYellingQuestion(remark) {
		return "Calm down, I know what I'm doing!"
	}

	return "Whatever."
}

func isQuestion(s string) bool {
	return strings.HasSuffix(s, "?")
}

func isYellingQuestion(s string) bool {
	s, _ = regexp.MatchString("[A-Z]", s)
	return s && strings.HasSuffix(s, "?")

}

func isYelling(s string) bool {
	hasUpperCase, _ := regexp.MatchString("[A-Z]{2,}", s)
	hasLowerCase, _ := regexp.MatchString("[a-z]+", s)
	return hasUpperCase && !hasLowerCase
}
