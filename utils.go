package main

import (
	"strings"
)

func ToCamelCase(word string) string {
	chunks := strings.Split(word, "_")

	for key, val := range(chunks) {
		chunks[key] = strings.Title(strings.ToLower(val))
	}

	return strings.Join(chunks, "")
}
