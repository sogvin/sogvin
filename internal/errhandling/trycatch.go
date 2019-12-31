package errhandling

import (
	"fmt"
	"io"
	"os"
)

func CopyFile_trycatch(src, dst string) (err error) {
	defer func() {
		e := recover()
		if e != nil {
			os.Remove(dst)
			err = fmt.Errorf("copy %s %s: %v", src, dst, err)
		}
	}()

	r := mustOpen(src)
	defer r.Close()
	w := mustCreate(dst)
	mustCopy(w, r)
	mustClose(w)
	return
}

func mustOpen(filename string) *os.File {
	fh, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return fh
}

func mustCreate(filename string) *os.File {
	fh, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	return fh
}

func mustCopy(w io.WriteCloser, r io.Reader) int64 {
	n, err := io.Copy(w, r)
	if err != nil {
		w.Close()
		panic(err)
	}
	return n
}

func mustClose(fh io.Closer) {
	err := fh.Close()
	if err != nil {
		panic(err)
	}
}
