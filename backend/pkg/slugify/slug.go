package slugify

import (
	"regexp"
	"strings"
)

func Slugify(text string) string {
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)

	result := strings.ToLower(text)

	result = strings.TrimSpace(result)

	result = reg.ReplaceAllString(result, "-")

	return result
}
