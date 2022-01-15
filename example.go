package sogvin

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	. "github.com/gregoryv/web"
)

func linkDrill(filename string) *Element {
	title := drillTitle(filename)
	return Li(A(Href(toHtmlFile(filename)), title))
}

func drillTitle(filename string) string {
	line := firstLine(filename)
	title := line
	parts := strings.Split(line, ";")
	if len(parts) > 1 {
		title = parts[1]
	}
	return title
}

func toHtmlFile(filename string) string {
	return strings.Replace(filename, ".go", ".html", 1)
}

func firstLine(filename string) string {
	fh := openFile(filename)
	defer fh.Close()
	line := readLine(fh)
	if line[:2] != "//" {
		log.Fatal("missing file comment: ", filename)
	}
	return line[3:] // skip first comment '// '
}

func openFile(filename string) io.ReadCloser {
	fh, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return fh
}

func readLine(r io.Reader) string {
	s := bufio.NewScanner(r)
	s.Scan()
	return s.Text()
}

// runExample first file contains init() that is renamed to main()
func runExample(args string, files ...string) ([]byte, error) {
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

	parts := strings.Split(args, " ")
	fullArgs := append(
		[]string{"run", filepath.Base(scriptFile)}, parts...,
	)
	cmd := exec.Command("go", fullArgs...)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()

	var buf bytes.Buffer
	buf.WriteString("$ ")
	buf.WriteString(cmd.String())
	buf.WriteString("\n")
	buf.Write(out)
	return buf.Bytes(), err
}
