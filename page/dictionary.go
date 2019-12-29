package page

import . "github.com/gregoryv/web/doctype"

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
