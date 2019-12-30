package notes

import (
	. "github.com/gregoryv/web/doctype"
)

var InlineTestHelpers = Html(
	Head(utf8, viewport, theme, a4),
	Body(
		header("", "Testing"),
		Article(
			H1("Inline test helpers"),
			P(
				"Use inline test helpers to minimize indentation and have",
				"failures point out failed cases directly.",
			),

			boxnote("Inlined helper does not need t argument.", 0.8),
			boxnote("Descriptive cases fail on correct line.", 5.6),
			loadGoFile("./testing/inline_test.go", 8, -1),

			boxnote("Utmost 2 inlined helpers.", 0.2),

			P(`Keep it simple and use utmost two inlined
	           helpers. Compared to table-driven-tests inlined helpers
	           declare the <em>how</em> before the cases.  If you have
	           many cases, this style is more readable as you first
	           tell the reader the meaning of &#34;ok&#34; and
	           &#34;bad&#34;.  <br> Another positive benefit of this
	           style is values are not grouped in a testcase
	           variable. I.e. readability improves as the values are
	           used directly.  <br>This style may be less readable if
	           each case requires many values, though it depends on
	           the lenght of the values combined.`,
			),
		),
		footer,
	),
)
