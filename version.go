package sogvin

import (
	_ "embed"
	"strings"
)

func Version() string {
	from := strings.Index(changelog, "## [")
	to := strings.Index(changelog[from:], "]")
	return changelog[from+4 : from+to]
}

func Released() string {
	from := strings.Index(changelog, "## [")
	from += strings.Index(changelog[from:], "]")
	to := strings.Index(changelog[from:], "\n")
	return changelog[from+1 : from+to]
}

//go:embed changelog.md
var changelog string
