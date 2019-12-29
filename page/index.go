package page

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
