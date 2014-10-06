package main

import (
	"fmt"
	"strconv"
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
		ok("%s contains '%s'", subject, expected)
	} else {
		ng("%s doesn't contain '%s'", subject, expected)
	}
}

func assertEqual(subject string, actual string, expected string) {
	if actual == expected {
		ok("%s is '%s'", subject, expected)
	} else {
		ng("%s isn't '%s'", subject, expected)
	}
}

func assertPresent(subject string, slice []string, expected string) {
	if ContainSlice(slice, expected) {
		ok("%s has '%s'", subject, expected)
	} else {
		ng("%s doesn't have '%s'", subject, expected)
	}
}

func (player *Player) getAttribute(selector string, name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("Attribute name is not defined")
	}

	el, err := player.FindElement(selector)

	if err != nil {
		return "", err
	}

	value, err := el.GetAttribute(name)

	if err != nil {
		return "", err
	}

	return value, nil
}

func (player *Player) PlayTitleEqualAssert(action Action) error {
	expected := action["expected"]
	title, err := player.wd.Title()

	if err != nil {
		return err
	}

	assertEqual("title", title, expected)

	return nil
}

func (player *Player) PlayTextEqualAssert(action Action) error {
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

	assertEqual(selector+" text", text, expected)

	return nil
}

func (player *Player) PlayTextContainAssert(action Action) error {
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

	assertContain(selector+" text", text, expected)

	return nil
}

func (player *Player) PlayAttributeEqualAssert(action Action) error {
	selector := action["element"]
	attrName := action["name"]
	expected := action["expected"]

	value, err := player.getAttribute(selector, attrName)

	if err != nil {
		return err
	}

	assertEqual(selector+"["+attrName+"]", value, expected)

	return nil
}

func (player *Player) PlayAttributeContainAssert(action Action) error {
	selector := action["element"]
	attrName := action["name"]
	expected := action["expected"]

	value, err := player.getAttribute(selector, attrName)

	if err != nil {
		return err
	}

	assertContain(selector+"["+attrName+"]", value, expected)

	return nil
}

func (player *Player) PlayAttributePresentAssert(action Action) error {
	selector := action["element"]
	attrName := action["name"]
	expected := action["expected"]

	value, err := player.getAttribute(selector, attrName)

	if err != nil {
		return err
	}

	assertPresent(selector+"["+attrName+"]", strings.Fields(value), expected)

	return nil
}

func (player *Player) PlayValueContainAssert(action Action) error {
	selector := action["element"]
	expected := action["expected"]

	value, err := player.getAttribute(selector, "value")

	if err != nil {
		return err
	}

	assertContain(selector+" value", value, expected)

	return nil
}

func (player *Player) PlayValueEqualAssert(action Action) error {
	selector := action["element"]
	expected := action["expected"]

	value, err := player.getAttribute(selector, "value")

	if err != nil {
		return err
	}

	assertEqual(selector+" value", value, expected)

	return nil
}

func (player *Player) PlayCssAssert(action Action) error {
	selector := action["element"]
	propertyName := action["property"]
	expected := action["expected"]

	el, err := player.FindElement(selector)

	if err != nil {
		return err
	}

	value, err := el.CSSProperty(propertyName)

	if err != nil {
		return err
	}

	assertEqual(selector+" "+propertyName, value, expected)

	return nil
}

func (player *Player) PlayLengthAssert(action Action) error {
	selector := action["element"]
	expected := action["expected"]

	els, err := player.FindElements(selector)

	if err != nil {
		return err
	}

	assertEqual(selector+" length", strconv.Itoa(len(els)), expected)

	return nil
}
