package errhandling

import (
	"fmt"
	"io"
	"os"
)

func CopyFile_split(src, dst string) (err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = fmt.Errorf("copy %s %s: %v", src, dst, e)
		}
	}()

	r := mustOpen(src)
	defer r.Close()
	w := mustCreate(dst)
	return copyFileStreams(w, r)
}

// copyFileStreams copies reader to writer and closes the writer
// On error, dst file is removed
func copyFileStreams(w NamedWriteCloser, r NamedReader) error {
	_, copyErr := io.Copy(w, r)
	closeErr := w.Close()

	if err := firstOf(copyErr, closeErr); err != nil {
		dst, src := w.Name(), r.Name()
		os.Remove(dst)
		return fmt.Errorf("copyFileStreams %s %s: %v", src, dst, err)
	}
	return nil
}

func firstOf(errors ...error) error {
	for _, err := range errors {
		if err != nil {
			return err
		}
	}
	return nil
}

type NamedWriteCloser interface {
	io.WriteCloser
	Name() string
}

type NamedReader interface {
	io.Reader
	Name() string
}
