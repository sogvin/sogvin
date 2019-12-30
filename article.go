package notes

import . "github.com/gregoryv/web/doctype"

var Index = Html(en,
	Head(utf8, viewport, theme, a4,
		Style(Type("text/css"),
			`ul {
      padding: 0px;
      }
      li {
      padding: 0px;
      list-style: none;
      }`),
	),
	Body(Article(H1("Software Engineering"),
		`<p>Notes by Gregory Vin&ccaron;i&cacute;</p>

      <h2>Table of Contents</h2>
      <ul>
	<li><a href="purpose_of_func_main.html">Purpose of func main()</a></li>
	<li><a href="nexus_pattern.html">Nexus pattern</a></li>
	<li><a href="inline_test_helpers.html">Inline test helpers</a></li>
	<li><a href="graceful_server_shutdown.html">Graceful server shutdown</a></li>
      </ul>`),
		footer,
	),
)

var Dictionary = Html(en,
	Head(utf8, viewport, theme, a4),
	Body(`<header>
      <span class="left"></span>
      <code>References</code>
    </header>

    <article>
      <h1>Dictionary</h1>

      <p>Short list of words/terms often used in software engineering
	and sometimes defined differently. Only domain agnostic terms
	have been listen, for the rest consult an english dictionary.
	I often use the <code>dict</code> command line tool.</p>

      <dl>
	<dt>Argument</dt>
	<dd>String following the command on the command line.</dd>

	<dt>Flag</dt>
	<dd>Boolean option.</dd>

	<dt>Option</dt>
	<dd>Argument starting with single or double dashes.</dd>

      </dl>
    </article>`,
		footer,
	),
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
			loadGoFile("./errhandling/nexus.go", 21, -1),

			`With the fileIO nexus inplace the CopyFile function is
	readable and with only one error checking and handling needed.`,
			loadGoFile("./errhandling/nexus.go", 8, 19),
		),
		footer,
	),
)
