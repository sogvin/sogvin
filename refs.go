package website

import (
	. "github.com/gregoryv/web"
)

func packageRefs() *Element {
	return Article(
		H1("Packages"),

		Ul(
			gregoryv("stamp", "build information code generator"),
			gregoryv("find", "files by name or content"),

			gregoryv("golden", "simplify use of golden files"),
			gregoryv("qual", "quality constraints"),
			gregoryv("ex", "indented JSON or redirect handler response to stdout"),
			gregoryv("uncover", "paths that need more testing"),

			gregoryv("draw", "software engineering diagrams"),
			gregoryv("web", "html generation"),
		),
	)
}

func github(pth, label string) *Element {
	return A(
		Href("https://github.com/"+pth),
		label,
	)
}
