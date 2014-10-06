package main

import (
	"path/filepath"
	"regexp"
	"strings"
)

var reHttpProto = regexp.MustCompile("^https?://")
var reAbsFilePath = regexp.MustCompile("^/")

func ToCamelCase(word string) string {
	chunks := strings.Split(word, "_")

	for key, val := range chunks {
		chunks[key] = strings.Title(strings.ToLower(val))
	}

	return strings.Join(chunks, "")
}

func NormalizeUrl(url string, baseDir string) string {
	if url == "" {
		return ""
	}

	if reHttpProto.MatchString(url) {
		return url
	}

	if reAbsFilePath.MatchString(url) {
		return "file://" + url
	}

	return "file://" + filepath.Join(baseDir, url)
}
