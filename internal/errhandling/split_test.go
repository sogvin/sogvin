package errhandling

import (
	"fmt"
	"os"
	"testing"
)

func Test_copyFileStreams_err_on_copy(t *testing.T) {
	wd, assert, cleanup := setupSrc(t)
	defer cleanup()

	r, _ := os.Open(wd.Join(src))
	err := copyFileStreams(&BadWriter{writeErr: fmt.Errorf("write")}, r)
	assert(err != nil).Error(err)
}

func Test_copyFileStreams_err_on_close(t *testing.T) {
	wd, assert, cleanup := setupSrc(t)
	defer cleanup()

	r, _ := os.Open(wd.Join(src))
	err := copyFileStreams(&BadWriter{closeErr: fmt.Errorf("close")}, r)
	assert(err != nil).Error(err)
}

type BadWriter struct {
	writeErr error
	closeErr error
}

func (w *BadWriter) Name() string { return "test" }
func (w *BadWriter) Write(b []byte) (int, error) {
	if w.writeErr != nil {
		return 0, w.writeErr
	}
	return len(b), nil

}
func (w *BadWriter) Close() error { return w.closeErr }
