package spec

import (
	"fmt"
)

type errHandler func(v ...interface{})

func (me errHandler) Error(err error) error {
	if me == nil {
		return nil
	}
	me(err)
	return err
}

func (me errHandler) Errorf(format string, args ...interface{}) error {
	if me == nil {
		return nil
	}
	err := fmt.Errorf(format, args...)
	me(err)
	return err
}
