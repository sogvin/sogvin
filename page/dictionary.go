package page

import . "github.com/gregoryv/web/doctype"

var Dictionary = Html(Lang("en"),
	`<head>
    <meta charset="utf-8"/>
    <title></title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" type="text/css" href="theme.css">
    <link rel="stylesheet" type="text/css" href="a4.css">
  </head>
  <body>
    <header>
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
    </article>

    <footer>
      Gregory Vinčić
    </footer>
  </body>
`)
