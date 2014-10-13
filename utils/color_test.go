package utils

import "testing"

func TestColor(t *testing.T) {
	if Black("foo") != "\x1b[30mfoo\x1b[0m" {
		t.Errorf("Failed Black: %s", Black("foo"))
	}

	if Red("foo") != "\x1b[31mfoo\x1b[0m" {
		t.Errorf("Failed Red: %s", Red("foo"))
	}

	if Green("foo") != "\x1b[32mfoo\x1b[0m" {
		t.Errorf("Failed Green: %s", Green("foo"))
	}

	if Yellow("foo") != "\x1b[33mfoo\x1b[0m" {
		t.Errorf("Failed Yellow: %s", Yellow("foo"))
	}

	if Blue("foo") != "\x1b[34mfoo\x1b[0m" {
		t.Errorf("Failed Blue: %s", Blue("foo"))
	}

	if Magenta("foo") != "\x1b[35mfoo\x1b[0m" {
		t.Errorf("Failed Magenta: %s", Magenta("foo"))
	}

	if Cyan("foo") != "\x1b[36mfoo\x1b[0m" {
		t.Errorf("Failed Cyan: %s", Cyan("foo"))
	}

	if White("foo") != "\x1b[37mfoo\x1b[0m" {
		t.Errorf("Failed White: %s", White("foo"))
	}
}
