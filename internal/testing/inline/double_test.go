package inline

import (
	"testing"
)

func Test_double(t *testing.T) {
	ok := func(input, exp int) {
		t.Helper()
		got, err := double(input)
		if err != nil {
			t.Error(err)
		}
		if got != exp {
			t.Errorf("double(%v) returned %v, expected %v", input, got, exp)
		}
	}
	// cases
	ok(1, 2)
	ok(3, 6)
	ok(MAX, MAX)

	bad := func(input, exp int) {
		t.Helper()
		_, err := double(input)
		if err == nil {
			t.Errorf("double(%v) should fail", input)
		}
	}
	bad(-2, 4)
}
