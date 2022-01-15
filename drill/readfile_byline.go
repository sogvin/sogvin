// Read file line by line
//
// When dealing with large files.
package drill

import (
	"bufio"
	"log"
	"os"
)

func init() {
	fh, err := os.Open("readfile_byline.go")
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()

	s := bufio.NewScanner(fh)

	for s.Scan() {
		println(s.Text())
		// only read the first line as an example
		break
	}
}
