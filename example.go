package sogvin

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// runExample first file contains init() that is renamed to main()
func runExample(cmdline string, files ...string) ([]byte, error) {
	first := files[0]

	data, err := ioutil.ReadFile(first)
	if err != nil {
		return nil, err
	}

	data = bytes.ReplaceAll(data, []byte("func init("), []byte("func main("))
	data = bytes.ReplaceAll(data, []byte("package drill"), []byte("package main"))

	// Use name of first file as command name, so we can have many
	// files in same directory, but speed up builds
	name := filepath.Base(first)
	i := strings.Index(name, ".")
	dir := filepath.Join(os.TempDir(), name[:i])
	if err := os.MkdirAll(dir, 0722); err != nil {
		return nil, err
	}
	log.Println(dir)

	scriptFile := filepath.Join(dir, name)
	if err := ioutil.WriteFile(scriptFile, data, 0644); err != nil {
		return nil, err
	}

	cmd := exec.Command("go", "run", filepath.Base(scriptFile))
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()

	var buf bytes.Buffer
	buf.WriteString("$ ")
	buf.WriteString(cmdline)
	buf.WriteString("\n")
	buf.Write(out)
	return buf.Bytes(), err
}
