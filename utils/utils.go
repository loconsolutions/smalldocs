package utils

import (
	"regexp"
	"strings"
)

// Regexp to remove special character
var SLUGIFY = regexp.MustCompile(`[^a-zA-Z0-9\.-]+`)
var TITLE = regexp.MustCompile(`[^a-zA-Z0-9 \.-]+`)

//
// Get formatted name from string
//
func Slug(str string) string {
	str = SLUGIFY.ReplaceAllString(str, " ")
	str = strings.ToLower(strings.Trim(str, " "))
	str = strings.Replace(str, " ", "-", -1)
	return str
}

//
// Formate title
//
func Title(str string) string {
	str = TITLE.ReplaceAllString(str, "")
	return strings.Trim(str, " ")
}
