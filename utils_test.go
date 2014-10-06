package main

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
