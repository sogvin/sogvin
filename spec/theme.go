package spec

import . "github.com/gregoryv/web"

func Theme() *CSS {
	css := NewCSS()
	css.Style("html, body",
		"margin: 0 0",
		"body: 0 0",
	)
	css.Style("body",
		"padding: 1em 1.618em 1em 1.618em",
		"max-width: 21cm",
	)
	css.Style("h1:first-child",
		"margin-top: 0",
	)
	css.Style("nav ul",
		"list-style-type: none",
		"padding-left: 0",
	)
	css.Style("li.h3",
		"margin-left: 1.618em",
	)
	return css
}
