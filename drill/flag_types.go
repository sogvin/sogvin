// Parse builtin types
//
// Convert options to correct type early.
package drill

import (
	"flag"
	"time"
)

func init() {
	var (
		n = flag.Int("i", 7, "integer")
		s = flag.String("s", "hi", "string")
		b = flag.Bool("b", false, "bool")
		d = flag.Duration("d", time.Second, "time.Duration")
	)
	flag.Parse()
	println(*n, *s, *b, *d)
}
