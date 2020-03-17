package internal

import (
	"io"
	"testing"
)

func TestAssert(t *testing.T) {
	ok, bad := Assert(t)
	ok(nil)
	bad(io.EOF)

	n := &noopT{}
	ok, _ = Assert(n)
	ok(io.EOF)
	n.AssertError(t)

	n = &noopT{}
	ok, _ = Assert(n)
	ok(io.EOF, "my error")
	n.AssertError(t)

	n = &noopT{}
	_, bad = Assert(n)
	bad(nil)
	n.AssertError(t)

	n = &noopT{}
	_, bad = Assert(n)
	bad(nil, "my error")
	n.AssertError(t)
}

type noopT struct {
	errorCalled bool
}

func (*noopT) Helper()                     {}
func (n *noopT) Error(args ...interface{}) { n.errorCalled = true }

func (n *noopT) AssertError(t *testing.T) {
	t.Helper()
	if !n.errorCalled {
		t.Error("t.Error never called")
	}
}
