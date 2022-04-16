package website

import (
	"fmt"

	. "github.com/gregoryv/web"
)

func embedVersionAndRevision() *Element {
	return Article(
		H1("Embed version and revision"),

		P(`When is this valuable? When publishing software for
       traceability and referenc.  I.e. for bug reports or
       documentation reference. Your applications can use flags such
       as -v or -version for this purpose. One way to modify variables
       during the build is via -ldflags.`),

		H2("Using -ldflags"),
		P("First declare a variable, not constant, in the main package."),
		loadFile("./internal/cmd/embedversion/main.go", 9, -1),

		P(`Then compile and change the version with`),

		shellCommand(

			`go build -ldflags "-X main.version=0.1" ./cmd/app`,
		),

		P(`You can also change multiple values in this way, let's add
		the revision as well`),

		shellCommand(

			func() string {
				return fmt.Sprint("go build ",
					`-ldflags "-X main.version=0.1 -X main.revision=alpha" `,
					"./cmd/app",
				)
			}(),
		),
		//
	)
}
