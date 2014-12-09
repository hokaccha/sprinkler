package action

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/hokaccha/sprinkler/utils"
	"github.com/sourcegraph/go-selenium"
)

func actionLog(name string, message string, args ...interface{}) {
	utils.Debug("Run action "+name+" - "+message, args...)
}

type ActionRunner interface {
	Run(interface{}) error
	New(*ActionOpts)
	GetResult() *ActionResult
}

type ActionOpts struct {
	Wd      selenium.WebDriver
	Name    string
	BaseDir string
	Params  interface{}
}

type ActionResult struct {
	IsAssert  bool
	Successed bool
	Message   string
}

type ActionRunnerGetter func() ActionRunner

var actionMap = map[string]ActionRunnerGetter{}

func RegisterAction(name string, fn ActionRunnerGetter) {
	actionMap[name] = fn
}

func RunAction(opts *ActionOpts) (*ActionResult, error) {
	runnerFn, ok := actionMap[opts.Name]

	if !ok {
		return nil, fmt.Errorf("unknown action: %s", opts.Name)
	}

	runner := runnerFn()
	runner.New(opts)
	err := runner.Run(opts.Params)

	if err != nil {
		return nil, err
	}

	return runner.GetResult(), nil
}

type ActionBase struct {
	Wd      selenium.WebDriver
	Name    string
	Result  *ActionResult
	BaseDir string
}

func (a *ActionBase) New(opts *ActionOpts) {
	a.Wd = opts.Wd
	a.Name = opts.Name
	a.BaseDir = opts.BaseDir
	a.Result = new(ActionResult)
}

func (a *ActionBase) GetResult() *ActionResult {
	return a.Result
}

func (a *ActionBase) parseParams(data interface{}, v interface{}) error {
	var err error
	rv := reflect.ValueOf(v)
	vv := rv.Elem()
	vt := vv.Type()

	if rv.Kind() != reflect.Ptr || vv.Kind() != reflect.Struct {
		panic("invalid argument")
	}

	switch data.(type) {
	case map[interface{}]interface{}:
		m := map[string]interface{}{}

		for key, val := range data.(map[interface{}]interface{}) {
			m[key.(string)] = val
		}

		for i := 0; i < vt.NumField(); i++ {
			structField := vt.Field(i)
			structValue := vv.Field(i)
			val, ok := m[structField.Tag.Get("name")]

			if !ok {
				continue
			}

			err = setValue(structField, structValue, reflect.ValueOf(val))

			if err != nil {
				return err
			}
		}
	default:
		structField := vt.Field(0)
		structValue := vv.Field(0)
		err = setValue(structField, structValue, reflect.ValueOf(data))

		if err != nil {
			return err
		}
	}

	return nil
}

func setValue(structField reflect.StructField, structValue, val reflect.Value) error {
	errFmt := "invalid parameter: %s must be %s"
	name := structField.Tag.Get("name")

	if structValue.Kind() == reflect.Ptr {
		e := structValue.Type().Elem()
		if e.Kind() != val.Kind() {
			return fmt.Errorf(errFmt, name, e.Kind())
		}
		elem := reflect.New(e)
		elem.Elem().Set(val)
		structValue.Set(elem)
	} else {
		if structValue.Kind() != val.Kind() {
			return fmt.Errorf(errFmt, name, structValue.Kind())
		}
		structValue.Set(val)
	}

	return nil
}

func (a *ActionBase) findElement(selector string) (selenium.WebElement, error) {
	if selector == "" {
		return nil, fmt.Errorf("selector is not defined")
	}

	return a.Wd.FindElement(selenium.ByCSSSelector, selector)
}

func (a *ActionBase) findElements(selector string) ([]selenium.WebElement, error) {
	if selector == "" {
		return nil, fmt.Errorf("selector is not defined")
	}

	return a.Wd.FindElements(selenium.ByCSSSelector, selector)
}

func (a *ActionBase) getAttribute(selector, name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("Attribute name is not defined")
	}

	el, err := a.findElement(selector)

	if err != nil {
		return "", err
	}

	value, err := el.GetAttribute(name)

	if err != nil {
		return "", err
	}

	return value, nil
}

func (a *ActionBase) assertUntil(timeout int, fn func() error) error {
	// TODO: Use goroutine
	interval := 200
	duration := time.Duration(interval) * time.Millisecond
	totalTime := 0

	wait := func() {
		totalTime += interval
		time.Sleep(duration)
	}

	for {
		err := fn()

		if a.Result.Successed == true {
			return nil
		}

		if totalTime < timeout {
			wait()
			continue
		}

		return err
	}
}

func (a *ActionBase) assertEqual(subject, actual, expected string) {
	message := fmt.Sprintf("%s is '%s'", subject, expected)
	ok := actual == expected

	if !ok {
		message = errMsg(message, actual)
	}

	a.assert(ok, message)
}

func (a *ActionBase) assertContain(subject, actual, expected string) {
	message := fmt.Sprintf("%s contains '%s'", subject, expected)
	ok := strings.Contains(actual, expected)

	if !ok {
		message = errMsg(message, actual)
	}

	a.assert(ok, message)
}

func (a *ActionBase) assertPresent(subject, actual, expected string) {
	message := fmt.Sprintf("%s has '%s'", subject, expected)
	ok := utils.IsContained(strings.Fields(actual), expected)

	if !ok {
		message = errMsg(message, actual)
	}

	a.assert(ok, message)
}

func (a *ActionBase) assert(ok bool, message string) {
	a.Result.IsAssert = true

	if ok {
		a.Result.Successed = true
		a.Result.Message = fmt.Sprintf("%s - %s", utils.Green("OK"), message)
	} else {
		a.Result.Message = fmt.Sprintf("%s - %s", utils.Red("NG"), message)
	}
}

func errMsg(message, actual string) string {
	return fmt.Sprintf("%s - %s: %s", message, utils.Yellow("Actual"), actual)
}
