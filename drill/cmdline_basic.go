// Parse options using cmdline.BasicParser
//
// Use package cmdline if you want control of single and double dash
// options.
package drill

import (
	"github.com/gregoryv/cmdline"
)

func init() {
	var (
		cli     = cmdline.NewBasicParser()
		verbose = cli.Flag("-v, --verbose")
	)
	cli.Parse()
	print(verbose)
}
