package main

import (
	"fmt"
	"strings"
)

func ok(message string, args ...interface{}) {
	fmt.Printf("\033[32mOK\033[0m - "+message+"\n", args...)
}

func ng(message string, args ...interface{}) {
	fmt.Printf("\033[31mNG\033[0m - "+message+"\n", args...)
}

func assertContain(subject string, actual string, expected string) {
	if strings.Contains(actual, expected) {
		ok("%s text contains '%s'", subject, expected)
	} else {
		ng("%s text doesn't contain '%s'", subject, expected)
	}
}

func assertEqual(subject string, actual string, expected string) {
	if actual == expected {
		ok("%s text is '%s'", subject, expected)
	} else {
		ng("%s text isn't '%s'", subject, expected)
	}
}

func (player *Player) PlayEqualTitleAssert(action Action) error {
	expected := action["expected"]
	title, err := player.wd.Title()

	if err != nil {
		return err
	}

	assertEqual("title", title, expected)

	return nil
}

func (player *Player) PlayEqualTextAssert(action Action) error {
	selector := action["element"]
	expected := action["expected"]

	el, err := player.FindElement(selector)

	if err != nil {
		return err
	}

	text, err := el.Text()

	if err != nil {
		return err
	}

	assertEqual(selector, text, expected)

	return nil
}

func (player *Player) PlayContainTextAssert(action Action) error {
	selector := action["element"]
	expected := action["expected"]

	el, err := player.FindElement(selector)

	if err != nil {
		return err
	}

	text, err := el.Text()

	if err != nil {
		return err
	}

	assertContain(selector, text, expected)

	return nil
}

func (player *Player) PlayAttributeAssert(action Action) error {
	selector := action["element"]
	attrName := action["name"]
	expected := action["expected"]

	if attrName == "" {
		return fmt.Errorf("Attribute name is not defined")
	}

	el, err := player.FindElement(selector)

	if err != nil {
		return err
	}

	value, err := el.GetAttribute(attrName)

	if err != nil {
		return err
	}

	assertEqual(selector + "[" + attrName + "]", value, expected)

	return nil
}

func (player *Player) PlayContainValueAssert(action Action) error {
	selector := action["element"]
	expected := action["expected"]

	el, err := player.FindElement(selector)

	if err != nil {
		return err
	}

	value, err := el.GetAttribute("value")

	if err != nil {
		return err
	}

	assertContain(selector + " value", value, expected)

	return nil
}

func (player *Player) PlayEqualValueAssert(action Action) error {
	selector := action["element"]
	expected := action["expected"]

	el, err := player.FindElement(selector)

	if err != nil {
		return err
	}

	value, err := el.GetAttribute("value")

	if err != nil {
		return err
	}

	assertEqual(selector + " value", value, expected)

	return nil
}
