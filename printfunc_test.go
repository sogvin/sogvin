package notes_test

import (
	"bytes"
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/gregoryv/notes"
)

func Test(t *testing.T) {
	ok := func(funcName, exp string) {
		t.Helper()
		var w bytes.Buffer
		err := notes.PrintFunc("printfunc_test.go", funcName, &w)
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
	notes.MustLoadFunc("printfunc_test.go", "c")

	err := notes.PrintFunc("x", "y", nil)
	if err == nil {
		t.Error("Should fail")
	}

	err = notes.PrintFunc("printfunc_test.go", "y", nil)
	if err == nil {
		t.Error("Should fail")
	}
	defer func() {
		e := recover()
		if e == nil {
			t.Fail()
		}
	}()
	notes.MustLoadFunc("x", "y")
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
