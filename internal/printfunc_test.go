package internal

import (
	"bytes"
	"testing"

	"github.com/gregoryv/asserter"
)

func Test(t *testing.T) {
	ok := func(funcName, exp string) {
		t.Helper()
		var w bytes.Buffer
		err := PrintFunc("printfunc_test.go", funcName, &w)
		if err != nil {
			t.Error(err)
		}
		assert := asserter.New(t)
		assert().Equals(w.String(), exp)
	}
	ok("Model", `func (c *Car) Model() string {
	return "tesla"
}`)
	ok("a", "func a() {}")
	ok("b", `func b(x int, v ...interface{}) *testing.T {
	return nil
}`)
	ok("c", `func c(x int, v ...interface{}) {
}`)
	MustLoadFunc("printfunc_test.go", "c")

	err := PrintFunc("x", "y", nil)
	if err == nil {
		t.Error("Should fail")
	}

	err = PrintFunc("printfunc_test.go", "y", nil)
	if err == nil {
		t.Error("Should fail")
	}
	defer catchPanic(t)
	MustLoadFunc("x", "y")
}

func catchPanic(t *testing.T) {
	t.Helper()
	e := recover()
	if e == nil {
		t.Fail()
	}
}

func a() {}

type Car struct{}

func (c *Car) Model() string {
	return "tesla"
}

func b(x int, v ...interface{}) *testing.T {
	return nil
}

func c(x int, v ...interface{}) {
}
