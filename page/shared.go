package page

import . "github.com/gregoryv/web/doctype"

var (
	en       = Lang("en")
	utf8     = Meta(Charset("utf-8"))
	viewport = Meta(
		Name("viewport"),
		Content("width=device-width, initial-scale=1.0"),
	)
	theme = stylesheet("theme.css")
	a4    = stylesheet("a4.css")
)

func stylesheet(href string) *Tag {
	return Link(
		Rel("stylesheet"),
		Type("text/css"),
		Href(href),
	)
}
