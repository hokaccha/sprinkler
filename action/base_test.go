package action

import "testing"

type T struct {
	Foo string  "foo"
	Bar int     "bar"
	Baz *string "baz"
}

func TestParseParams(t *testing.T) {
	var err error
	a := &ActionBase{}

	var d interface{} = map[interface{}]interface{}{
		"foo": "foo val",
		"bar": 100,
		"baz": "baz val",
	}

	t1 := &T{}

	err = a.parseParams(d, t1)

	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	if t1.Foo != "foo val" {
		t.Errorf(`t1.Foo should be "foo val"`)
	}

	if t1.Bar != 100 {
		t.Errorf("t1.Bar should be 100")
	}

	if *t1.Baz != "baz val" {
		t.Errorf(`*t1.Baz should be "baz val"`)
	}

	t2 := &T{}
	err = a.parseParams("foo val2", t2)

	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	if t2.Foo != "foo val2" {
		t.Errorf(`t2.Foo should be "foo val2"`)
	}

	if t2.Bar != 0 {
		t.Errorf("t2.Bar should be 0")
	}

	if t2.Baz != nil {
		t.Errorf(`t2.Baz should be nil`)
	}

	t3 := &T{}

	err = a.parseParams(map[interface{}]interface{}{"foo": 100}, t3)
	if err.Error() != "invalid parameter: foo must be string" {
		t.Errorf("invalid err: \"%s\"", err.Error())
	}

	err = a.parseParams(map[interface{}]interface{}{"bar": "foo"}, t3)
	if err.Error() != "invalid parameter: bar must be int" {
		t.Errorf("invalid err: \"%s\"", err.Error())
	}

	err = a.parseParams(map[interface{}]interface{}{"baz": 100}, t3)
	if err.Error() != "invalid parameter: baz must be string" {
		t.Errorf("invalid err: \"%s\"", err.Error())
	}

	err = a.parseParams(100, t3)
	if err.Error() != "invalid parameter: foo must be string" {
		t.Errorf("invalid err: \"%s\"", err.Error())
	}
}
