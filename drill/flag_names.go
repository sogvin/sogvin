// Short and long option names
//
// Simplify for users of your programs by providing both a long and a
// short option variation. This drill uses the builtin package flag.
package drill

import (
	"flag"
)

func init() {
	verbose, usage := false, ""
	flag.BoolVar(&verbose, "v", verbose, usage)
	flag.BoolVar(&verbose, "verbose", false, usage)

	flag.Parse()
}
