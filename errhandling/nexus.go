package errhandling

import (
	"io"
	"os"
)

func CopyFile_nexus(src, dst string) error {
	fi := &fileIO{}
	r := fi.Open(src)
	w := fi.Create(dst)
	fi.Copy(w, r)
	fi.Close(w)
	if fi.err != nil {
		os.Remove(dst)
	}
	return fi.err
}

type fileIO struct {
	err error
}

func (fi *fileIO) Open(filename string) (fh *os.File) {
	if fi.err != nil {
		return
	}
	fh, fi.err = os.Open(filename)
	return
}

func (fi *fileIO) Create(filename string) (fh *os.File) {
	if fi.err != nil {
		return
	}
	fh, fi.err = os.Create(filename)
	return
}

func (fi *fileIO) Close(w io.Closer) {
	if fi.err != nil {
		return
	}
	fi.err = w.Close()
	return
}

func (fi *fileIO) Copy(w io.WriteCloser, r io.Reader) (n int64) {
	if fi.err != nil {
		return
	}
	n, fi.err = io.Copy(w, r)
	if fi.err != nil {
		w.Close()
	}
	return
}
