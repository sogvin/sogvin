package errhandling

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_mustCopy(t *testing.T) {
	defer catch(t)
	w := &BadWriter{writeErr: fmt.Errorf("bad write")}
	r := bytes.NewBufferString("hello")
	mustCopy(w, r)
}

func Test_mustClose(t *testing.T) {
	defer catch(t)
	w := &BadWriter{closeErr: fmt.Errorf("bad close")}
	mustClose(w)
}

func catch(t *testing.T) {
	t.Helper()
	e := recover()
	if e == nil {
		t.Error("Didn't panic")
	}
}
