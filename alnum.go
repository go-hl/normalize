package normalize

import (
	"regexp"
	"strings"
)

// Alnum retain only single white spaces between words and alnum chars. At end trasnform to lower.
//
// E.g.:
//
//	` @h3LLo !|   W0#Rld  ` -> `h3llo w0rld`
func Alnum(str string) string {
	// exclude any punct char
	alnum := regexp.MustCompile(`[[:punct:]]`).ReplaceAllString(str, "")

	// get all chars execpt white spaces
	fields := strings.Fields(alnum)

	// joins the slice of chars with a single white spce
	join := strings.Join(fields, " ")

	// transform all to lower
	lower := strings.ToLower(join)

	return lower
}
