package errhandling

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/gregoryv/asserter"
)

func Test_writeAlog(t *testing.T) {
	buf := bytes.NewBufferString("")
	err := writeAlot(buf)
	if err != nil {
		t.Fail()
	}
}

func Test_fprint_nexus(t *testing.T) {
	err := fmt.Errorf("already failed")
	fprint := fprintFunc(err)
	n, got := fprint(ioutil.Discard, "Hello")
	assert := asserter.New(t)
	assert(got != nil).Error(got)
	assert(n == 0).Errorf("Should not have written anything, wrote: %v", n)
}
