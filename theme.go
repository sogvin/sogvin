package sogvin

import (
	. "github.com/gregoryv/web"
)

func A4() *CSS {
	css := NewCSS()
	css.Filename = "a4.css"

	css.Style("html, body, header, footer, h1, h2, h3, h4, h5",
		"margin: 0px 0px",
		"padding: 0px 0px",
		"background-color: #ffffff",
	)
	css.Style("body",
		"width: 21cm",
	)
	css.Style("header",
		"text-align: right",
	)
	css.Style("p",
		"line-height: 1.4em",
	)
	css.Style("span.left",
		"float: left",
	)
	css.Style("article",
		"margin-top: 0cm",
		"margin-bottom: 0.5cm",
		"padding-top: 0.2cm",
		"padding-left: 4cm",
	)
	css.Style("h1, h2",
		"margin-left: -4cm",
	)
	css.Style("h3",
		"margin-bottom:0.5cm",
	)
	css.Style("footer",
		"text-align: right",
	)
	css.Style("pre code",
		"font-size: 14px",
	)
	css.Style("h1 code, h2 code",
		"font-weight: normal",
		"margin-left: 8px",
	)
	css.Style("code.go",
		"-moz-tab-size: 4",
		"tab-size: 4",
	)
	css.Style(".command, .srcfile",
		"margin-top: 1.6em",
		"margin-bottom: 1.6em",
		"padding-left: 1.6em",
		"background-color: #eaeaea",
	)
	css.Style(".srcfile code",
		"padding: .6em 0 .6em 0",
		"background-position: right top",
		"display: block",
	)
	css.Style(".complete",
		"border: 1px solid #727272",
	)
	css.Style(".filename",
		"display: block",
		"text-align: right",
		"margin-bottom: -1.6em",
		"font-family: mono",
		"font-size: 12px",
	)
	css.Style(".command",
		"border-left: 7px #727272 solid",
		"padding: .6em 1.6em .6em 1.6em",
	)
	css.Style(".sidenote",
		"border: 1px solid black",
		"width: 3.3cm",
		"padding: 1px 1px",
		"font-size: 0.8em",
		"position: absolute",
		"margin-left: -4cm",
	)
	css.Style(".inner",
		"padding: 0.1cm",
		"border: 1px solid black",
	)
	css.Style("article.toc h3",
		"margin-left: 0",
	)
	css.Style("article.toc ul",
		"padding: 0px",
	)
	css.Style("article.toc ul li",
		"padding: 0px",
		"list-style: none",
	)
	css.Style("div.figure",
		"text-align: center",
	)

	screen := css.Media("screen")
	screen.Style("html, body",
		"margin: 3px 10px",
		"padding: 3px 10px",
	)

	print := css.Media("print")
	print.Style("footer",
		"position: absolute",
		"bottom: 0px",
		"left: 0px",
		"width: 100%",
	)
	return css
}

func Theme() *CSS {
	css := NewCSS()
	css.Filename = "theme.css"

	fonts := "Inconsolata|Source+Sans+Pro"
	css.Import("https://fonts.googleapis.com/css?family=" + fonts)

	css.Style("html, body",
		"font-family: 'Source Sans Pro', sans-serif",
	)
	css.Style("code",
		"font-family: Inconsolata",
	)
	css.Style("quote",
		"font-style: italic",
	)

	print := css.Media("print")
	print.Style("    a",
		"text-decoration: none",
		"color: black",
	)

	return css
}
