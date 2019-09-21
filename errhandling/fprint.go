package errhandling

import (
	"fmt"
	"io"
)

func writeAlot_original(w io.Writer) error {
	_, err := fmt.Fprint(w, "Hello")
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(w, "Sir")
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(w, "Lancelot")
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(w, "Where is your sword?")
	return err
}

func writeAlot(w io.Writer) (err error) {
	fprint := fprintFunc(err)
	fprint(w, "Hello")
	fprint(w, "Sir")
	fprint(w, "Lancelot")
	fprint(w, "Where is your sword?")
	return
}

type Fprinter func(io.Writer, ...interface{}) (int, error)

func fprintFunc(err error) Fprinter {
	return func(w io.Writer, args ...interface{}) (int, error) {
		if err != nil {
			return 0, err
		}
		var n int
		n, err = fmt.Fprint(w, args...)
		return n, err
	}
}
