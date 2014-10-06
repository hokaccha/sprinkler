package main

import (
	"fmt"
	"strconv"
	"strings"
)

func showOK(message string) {
	fmt.Printf("\033[32mOK\033[0m - %s\n", message)
}

func showNG(message string, actual string) {
	fmt.Printf("\033[31mNG\033[0m - %s - \033[33mActual\033[0m: %s\n", message, actual)
}

func test(ok bool, actual, message string) {
	if ok {
		showOK(message)
	} else {
		showNG(message, actual)
	}
}

func assertEqual(subject, actual, expected string) {
	message := fmt.Sprintf("%s is '%s'", subject, expected)
	ok := actual == expected

	test(ok, actual, message)
}

func assertContain(subject, actual, expected string) {
	message := fmt.Sprintf("%s contains '%s'", subject, expected)
	ok := strings.Contains(actual, expected)

	test(ok, actual, message)
}

func assertPresent(subject string, slice []string, expected string) {
	message := fmt.Sprintf("%s has '%s'", subject, expected)
	ok := ContainSlice(slice, expected)

	test(ok, fmt.Sprintf("%s", slice), message)
}

func (player *Player) getAttribute(selector, name string) (string, error) {
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

func (player *Player) PlayElementExistAssert(action Action) error {
	selector := action["element"]

	_, err := player.FindElement(selector)

	// TODO: See status code
	ok := err == nil
	message := fmt.Sprintf("%s exists", selector)
	test(ok, strconv.FormatBool(ok), message)

	return nil
}
