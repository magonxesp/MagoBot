package domain

import (
	"log/slog"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func RemoveAccents(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, err := transform.String(t, s)
	if err != nil {
		slog.Warn("error removing accent", "error", err)
		return s
	}

	return result
}

func Slugify(str string) string {
	regex := regexp.MustCompile(`[^A-Za-z0-9.]+`)
	result := RemoveAccents(str)
	result = strings.ToLower(result)
	result = strings.Trim(result, "")
	result = regex.ReplaceAllString(result, "_")
	return result
}
