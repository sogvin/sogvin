package main

import (
	"testing"

	"github.com/gregoryv/asserter"
)

func TestStarCounter(t *testing.T) {
	ok, bad := asserter.NewErrors(t)

	ok(NewStarCounter("name").Run())
	ok(NewStarCounter("name", "-h").Run())
	ok(NewStarCounter("name", "-help").Run())
	ok(NewStarCounter("name", "-size", "small").Run())
	ok(NewStarCounter("name", "-weight", "heavy").Run())

	bad(NewStarCounter("", "-nosuch").Run())
}
