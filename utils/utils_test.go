package utils

import (
	"testing"
)

func TestToCamelCase(t *testing.T) {
	test := func(input string, output string) {
		if ToCamelCase(input) != output {
			t.Errorf("Camelize %s is not %s: %s", input, output, ToCamelCase(input))
		}
	}

	test("foo", "Foo")
	test("foo_bar", "FooBar")
	test("foo_bar_baz", "FooBarBaz")
	test("foo_BAR_bAz", "FooBarBaz")
}

func TestNormalizeUrl(t *testing.T) {
	test := func(url string, baseDir string, output string) {
		if NormalizeUrl(url, baseDir) != output {
			t.Errorf("NormalizeUrl(\"%s\", \"%s\") is not %s: %s", url, baseDir, output, NormalizeUrl(url, baseDir))
		}
	}

	test("", "/path/to/dir", "")
	test("http://hoge", "", "http://hoge")
	test("https://hoge", "", "https://hoge")
	test("/foo/bar", "/path/to/dir", "file:///foo/bar")
	test("foo/bar", "/path/to/dir", "file:///path/to/dir/foo/bar")
	test("../foo/bar", "/path/to/dir", "file:///path/to/foo/bar")
}

func TestAbsPath(t *testing.T) {
	test := func(path string, baseDir string, output string) {
		if AbsPath(path, baseDir) != output {
			t.Errorf("AbsPath(\"%s\", \"%s\") is not %s: %s", path, baseDir, output, AbsPath(path, baseDir))
		}
	}

	test("/foo/bar", "/path/to/dir", "/foo/bar")
	test("foo/bar", "/path/to/dir", "/path/to/dir/foo/bar")
	test("../foo/bar", "/path/to/dir", "/path/to/foo/bar")
}

func TestContainSlice(t *testing.T) {
	if IsContained([]string{"foo", "bar", "baz"}, "baz") == false {
		t.Errorf("ContainSlice should return true")
	}

	if IsContained([]string{"foo", "bar", "baz"}, "hoge") == true {
		t.Errorf("ContainSlice should return false")
	}
}

func TestHasIntersection(t *testing.T) {
	test := func(a, b []string, expected bool) {
		if HasIntersection(a, b) != expected {
			t.Errorf("HasIntersection(%#v, %#v) should be %v", a, b, expected)
		}
	}

	test([]string{"foo", "bar"}, []string{"baz"}, false)
	test([]string{}, []string{}, false)
	test([]string{"foo", "bar"}, []string{}, false)
	test([]string{}, []string{"foo", "bar"}, false)

	test([]string{"foo", "bar"}, []string{"bar"}, true)
	test([]string{"foo"}, []string{"foo", "bar"}, true)
}

func TestStringSlice(t *testing.T) {
	var in interface{}

	test := func(input interface{}, expected []string) {
		m := "StringSlice(%#v) should be %#v"
		res := StringSlice(in)

		if len(res) != len(expected) {
			t.Errorf(m, in, expected)
		}

		for k, v := range StringSlice(in) {
			if expected[k] != v {
				t.Errorf(m, in, expected)
			}
		}
	}

	in = "foo"
	test(in, []string{"foo"})

	in = []string{"foo", "bar"}
	test(in, []string{"foo", "bar"})

	in = []interface{}{"foo", "bar"}
	test(in, []string{"foo", "bar"})

	in = nil
	test(in, nil)
}
