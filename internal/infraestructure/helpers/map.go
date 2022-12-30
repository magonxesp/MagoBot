package helpers

import "fmt"

func MapToKeyValueList(keyValues map[string]string, separator string) []string {
	var pairs []string

	for key, value := range keyValues {
		pairs = append(pairs, fmt.Sprintf("%s%s%s", key, separator, value))
	}

	return pairs
}
