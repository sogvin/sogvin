package errhandling

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_fileNexus_Open(t *testing.T) {
	x := &fileIO{err: fmt.Errorf("error")}
	fh := x.Open("anything")
	if fh != nil {
		t.Error("Open should not return anything when nexus has error")
	}
}

func Test_fileNexus_Copy(t *testing.T) {
	x := &fileIO{err: fmt.Errorf("error")}
	n := x.Copy(
		&testCloser{bytes.NewBufferString(""), nil},
		bytes.NewBufferString("hello"),
	)
	if n != 0 {
		t.Error("Copy should not copy anything when nexus has error")
	}
	x.err = nil
	x.Copy(
		&BadWriter{writeErr: fmt.Errorf("write error")},
		bytes.NewBufferString("hello"),
	)
	if x.err == nil {
		t.Error("Copy should set nex.err when closing fails")
	}
}

type testCloser struct {
	*bytes.Buffer
	err error
}

func (tc *testCloser) Close() error {
	return tc.err
}
