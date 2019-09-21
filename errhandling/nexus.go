package errhandling

import (
	"io"
	"os"
)

func CopyFile_nexus(src, dst string) error {
	x := &fileIO{}
	r := x.Open(src)
	w := x.Create(dst)
	x.Copy(w, r)
	x.Close(w)
	if x.err != nil {
		os.Remove(dst)
	}
	return x.err
}

type fileIO struct {
	err error
}

func (x *fileIO) Open(filename string) (fh *os.File) {
	if x.err != nil {
		return
	}
	fh, x.err = os.Open(filename)
	return
}

func (x *fileIO) Create(filename string) (fh *os.File) {
	if x.err != nil {
		return
	}
	fh, x.err = os.Create(filename)
	return
}

func (x *fileIO) Close(w io.Closer) {
	if x.err != nil {
		return
	}
	x.err = w.Close()
	return
}

func (x *fileIO) Copy(w io.WriteCloser, r io.Reader) (n int64) {
	if x.err != nil {
		return
	}
	n, x.err = io.Copy(w, r)
	if x.err != nil {
		w.Close()
	}
	return
}
