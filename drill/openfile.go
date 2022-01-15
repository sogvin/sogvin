// Open file for reading
package drill

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	fh, err := os.Open("openfile.go")
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()
	n, _ := io.Copy(ioutil.Discard, fh)
	print(n, " bytes")
}
