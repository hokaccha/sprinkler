package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/visionmedia/go-debug"
	"gopkg.in/yaml.v2"
)

var Debug = debug.Debug("sprinkler")

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

func IsContained(s []string, val string) bool {
	for _, v := range s {
		if val == v {
			return true
		}
	}

	return false
}

func ReadFile(path string) ([]byte, error) {
	reader, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer reader.Close()

	return ioutil.ReadAll(reader)
}

func LoadYAML(path string, out interface{}) error {
	data, err := ReadFile(path)

	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, out)
}

func MarshalYAML(m interface{}) (string, error) {
	b, err := yaml.Marshal(&m)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

func HasIntersection(a, b []string) bool {
	for _, v := range a {
		if IsContained(b, v) {
			return true
		}
	}

	return false
}

func StringSlice(in interface{}) []string {
	switch in.(type) {
	case string:
		return []string{in.(string)}
	case []string:
		return in.([]string)
	case []interface{}:
		ret := []string{}
		for _, v := range in.([]interface{}) {
			ret = append(ret, v.(string))
		}
		return ret
	case nil:
		return nil
	default:
		panic("invalid arguments")
	}

	return nil
}
