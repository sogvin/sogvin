package okbad

import (
	"testing"
)

func Test_double(t *testing.T) {
	var r int
	ok := assertOk(t)
	ok(double(&r, 1))
	ok(double(&r, 3))
	ok(double(&r, MAX))

	_k := assertBad(t)
	_k(double(&r, -2))
	_k(double(nil, 2))

	// verify data, some is good
	check := func(i, exp int) {
		t.Helper()
		var got int
		double(&got, i)
		if got != exp {
			t.Errorf("got %v, exp %v", got, exp)
		}
	}
	check(0, 0)     // edge
	check(1, 2)     // ok
	check(MAX, MAX) // other edge
	check(-1, 0)    // bad
}
