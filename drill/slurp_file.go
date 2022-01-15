// Slurp file
//
// Most of the time files are small and can easily be read all at once
// before doing something with it.
package drill

import (
	"log"
	"os"
)

func init() {
	// since go1.17, for older use ioutil.ReadFile
	data, err := os.ReadFile("slurp_file.go")
	if err != nil {
		log.Fatal(err)
	}
	print(len(data))
}
