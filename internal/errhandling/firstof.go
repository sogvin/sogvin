package errhandling

import (
	"io"
	"os"
)

func CopyFile_firstOf(src, dst string) error {
	var (
		r, openErr   = os.Open(src)
		w, createErr = os.Create(dst)
		_, copyErr   = io.Copy(w, r)
		closeErr     = w.Close()
		err          = firstOf(openErr, createErr, copyErr, closeErr)
	)
	if err != nil {
		os.Remove(dst)
	}
	return err
}
