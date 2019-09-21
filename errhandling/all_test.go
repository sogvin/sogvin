package errhandling

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/gregoryv/workdir"
)

const (
	src     = "from.txt"
	dest    = "to.txt"
	content = "hello"
)

func Test_implementations(t *testing.T) {
	cases := []struct {
		subtest string
		fn      copyFunc
	}{
		{"CopyFile try catch", copyFunc(CopyFile_trycatch)},
		{"Split CopyFile2", copyFunc(CopyFile_split)},
		{"Nexus", copyFunc(CopyFile_nexus)},
		{"firstOf", copyFunc(CopyFile_firstOf)},
		// Original cannot be covered
		// {"Original CopyFile", copyFunc(CopyFile)},
	}
	for _, c := range cases {
		t.Run(c.subtest, func(t *testing.T) {
			test_CopyFile_err_on_open(t, c.fn)
			test_CopyFile_err_on_create(t, c.fn)
			test_CopyFile_err_on_copy(t, c.fn)
			test_CopyFile_err_on_close(t, c.fn)
			test_CopyFile_ok_pathway(t, c.fn)
		})
	}
}

type copyFunc func(src, dst string) error

func (fn copyFunc) CopyFile(src, dst string) error {
	return fn(src, dst)
}

func test_CopyFile_err_on_open(t *testing.T, cp copyFunc) {
	t.Helper()
	wd, assert, cleanup := setupSrc(t)
	defer cleanup()

	// Make src unreadable
	os.Chmod(wd.Join(src), 0000)
	err := cp(wd.Join(src), wd.Join(dest))
	assert(err != nil).Error(err)
}

func test_CopyFile_err_on_create(t *testing.T, cp copyFunc) {
	t.Helper()
	wd, assert, cleanup := setupSrc(t)
	defer cleanup()

	// Make destination unwritable
	wd.Touch(dest)
	os.Chmod(wd.Join(dest), 0000)
	err := cp(wd.Join(src), wd.Join(dest))
	assert(err != nil).Error(err)
}

func test_CopyFile_err_on_copy(t *testing.T, cp copyFunc) {
	t.Helper()
	// impossible to test, signals that the CopyFile func does to much
}

func test_CopyFile_err_on_close(t *testing.T, cp copyFunc) {
	t.Helper()
	// impossible to test, signals that the CopyFile func does to much
}

func test_CopyFile_ok_pathway(t *testing.T, cp copyFunc) {
	t.Helper()
	wd, assert, cleanup := setupSrc(t)
	defer cleanup()

	err := cp(wd.Join(src), wd.Join(dest))
	assert(err == nil).Fatal(err)
	got, err := ioutil.ReadFile(wd.Join(dest))
	assert(err == nil).Fatal(err)
	assert().Equals(string(got), content)
}

func setupSrc(t *testing.T) (workdir.WorkDir, asserter.AssertFunc,
	func()) {
	t.Helper()
	assert := asserter.New(t)
	wd, tmpdirErr := workdir.TempDir()
	assert(tmpdirErr == nil).Fatal(tmpdirErr)
	writeErr := wd.WriteFile(src, []byte(content))
	assert(writeErr == nil).Fatal(writeErr)
	return wd, assert, func() {
		wd.RemoveAll()
		t.Log("cleanup")
	}
}
