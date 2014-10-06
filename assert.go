package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (player *Player) AssertEqual(subject, actual, expected string) {
	message := fmt.Sprintf("%s is '%s'", subject, expected)
	ok := actual == expected

	player.Test(ok, actual, message)
}

func (player *Player) AssertContain(subject, actual, expected string) {
	message := fmt.Sprintf("%s contains '%s'", subject, expected)
	ok := strings.Contains(actual, expected)

	player.Test(ok, actual, message)
}

func (player *Player) AssertPresent(subject string, slice []string, expected string) {
	message := fmt.Sprintf("%s has '%s'", subject, expected)
	ok := ContainSlice(slice, expected)

	player.Test(ok, fmt.Sprintf("%s", slice), message)
}

func (player *Player) Test(ok bool, actual, message string) {
	if ok {
		player.successCount++
		showOK(message)
	} else {
		player.failCount++
		showNG(message, actual)
	}
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
	expected := action["expect"]
	title, err := player.wd.Title()

	if err != nil {
		return err
	}

	player.AssertEqual("title", title, expected)

	return nil
}

func (player *Player) PlayTextEqualAssert(action Action) error {
	selector := action["element"]
	expected := action["expect"]

	el, err := player.FindElement(selector)

	if err != nil {
		return err
	}

	text, err := el.Text()

	if err != nil {
		return err
	}

	player.AssertEqual(selector+" text", text, expected)

	return nil
}

func (player *Player) PlayTextContainAssert(action Action) error {
	selector := action["element"]
	expected := action["expect"]

	if expected == "" {
		return fmt.Errorf("expect is not defined")
	}

	el, err := player.FindElement(selector)

	if err != nil {
		return err
	}

	text, err := el.Text()

	if err != nil {
		return err
	}

	player.AssertContain(selector+" text", text, expected)

	return nil
}

func (player *Player) PlayAttributeEqualAssert(action Action) error {
	selector := action["element"]
	attrName := action["name"]
	expected := action["expect"]

	value, err := player.getAttribute(selector, attrName)

	if err != nil {
		return err
	}

	player.AssertEqual(selector+"["+attrName+"]", value, expected)

	return nil
}

func (player *Player) PlayAttributeContainAssert(action Action) error {
	selector := action["element"]
	attrName := action["name"]
	expected := action["expect"]

	value, err := player.getAttribute(selector, attrName)

	if err != nil {
		return err
	}

	player.AssertContain(selector+"["+attrName+"]", value, expected)

	return nil
}

func (player *Player) PlayAttributePresentAssert(action Action) error {
	selector := action["element"]
	attrName := action["name"]
	expected := action["expect"]

	value, err := player.getAttribute(selector, attrName)

	if err != nil {
		return err
	}

	player.AssertPresent(selector+"["+attrName+"]", strings.Fields(value), expected)

	return nil
}

func (player *Player) PlayValueContainAssert(action Action) error {
	selector := action["element"]
	expected := action["expect"]

	value, err := player.getAttribute(selector, "value")

	if err != nil {
		return err
	}

	player.AssertContain(selector+" value", value, expected)

	return nil
}

func (player *Player) PlayValueEqualAssert(action Action) error {
	selector := action["element"]
	expected := action["expect"]

	value, err := player.getAttribute(selector, "value")

	if err != nil {
		return err
	}

	player.AssertEqual(selector+" value", value, expected)

	return nil
}

func (player *Player) PlayCssPropertyEqualAssert(action Action) error {
	selector := action["element"]
	propertyName := action["property"]
	expected := action["expect"]

	el, err := player.FindElement(selector)

	if err != nil {
		return err
	}

	value, err := el.CSSProperty(propertyName)

	if err != nil {
		return err
	}

	player.AssertEqual(selector+" "+propertyName, value, expected)

	return nil
}

func (player *Player) PlayElementLengthEqualAssert(action Action) error {
	selector := action["element"]
	expected := action["expect"]

	els, err := player.FindElements(selector)

	if err != nil {
		return err
	}

	player.AssertEqual(selector+" length", strconv.Itoa(len(els)), expected)

	return nil
}

func (player *Player) PlayElementExistAssert(action Action) error {
	selector := action["element"]

	_, err := player.FindElement(selector)

	// TODO: See status code
	ok := err == nil
	message := fmt.Sprintf("%s exists", selector)
	player.Test(ok, "not exists", message)

	return nil
}

func (player *Player) PlayElementVisibleAssert(action Action) error {
	selector := action["element"]

	el, err := player.FindElement(selector)

	if err != nil {
		return err
	}

	visible, err := el.IsDisplayed()

	if err != nil {
		return err
	}

	message := fmt.Sprintf("%s is visible", selector)
	player.Test(visible, "hidden", message)

	return nil
}

func (player *Player) PlayElementHiddenAssert(action Action) error {
	selector := action["element"]

	el, err := player.FindElement(selector)

	if err != nil {
		return err
	}

	visible, err := el.IsDisplayed()

	if err != nil {
		return err
	}

	message := fmt.Sprintf("%s is hidden", selector)
	player.Test(!visible, "visible", message)

	return nil
}

func (player *Player) PlayUrlEqualAssert(action Action) error {
	expected := action["expect"]

	url, err := player.wd.CurrentURL()

	if err != nil {
		return err
	}

	player.AssertEqual("url", url, expected)

	return nil
}

func (player *Player) PlayUrlContainAssert(action Action) error {
	expected := action["expect"]

	url, err := player.wd.CurrentURL()

	if err != nil {
		return err
	}

	player.AssertContain("url", url, expected)

	return nil
}

func (player *Player) PlayAlertTextEqualAssert(action Action) error {
	expected := action["expect"]

	text, err := player.wd.AlertText()

	if err != nil {
		return err
	}

	player.AssertEqual("alert text", text, expected)

	return nil
}

func (player *Player) PlayAlertTextContainAssert(action Action) error {
	expected := action["expect"]

	text, err := player.wd.AlertText()

	if err != nil {
		return err
	}

	player.AssertContain("alert text", text, expected)

	return nil
}

func showOK(message string) {
	fmt.Printf("\033[32mOK\033[0m - %s\n", message)
}

func showNG(message string, actual string) {
	fmt.Printf("\033[31mNG\033[0m - %s - \033[33mActual\033[0m: %s\n", message, actual)
}
