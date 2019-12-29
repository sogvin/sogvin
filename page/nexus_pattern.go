package page

import (
	. "github.com/gregoryv/web/doctype"
)

var NexusPattern = Html(en,
	Head(utf8, viewport, theme, a4),
	Body(
		header("", "Nexus pattern"),
		Article(
			H1("Nexus pattern"),
			P(
				"The word nexus is defined as",
				Quote(
					"&#34;The means of connection between things linked in series&#34;",
				),
				"The pattern is useful in",
				A(
					Href(
						"https://go.googlesource.com/proposal/+/master/design"+
							"/go2draft-error-handling-overview.md",
					),
					"error handling",
				),
				"sequential function calls.",
			),

			H2("Example <code>CopyFile(from, to string)</code>"),
			P(`Copying a file, if done all in one function, is unreadable due to
multiple error checking and handling.  With the nexus pattern you
define a <code>type fileIO struct</code> with the error field. Each
method must check the previous error and return if it is set without
doing anything. This way all subsequent calls are no-operations.`),

			boxnote("The err field links operations.", 0.6),
			boxnote("Each method sets x.err before returning.", 3.3),
			loadGoFile("../../../errhandling/nexus.go", 21, -1),

			`With the fileIO nexus inplace the CopyFile function is
	readable and with only one error checking and handling needed.`,
			loadGoFile("../../../errhandling/nexus.go", 8, 19),
		),
		footer,
	),
)
