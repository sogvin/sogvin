package website

import (
	. "github.com/gregoryv/web"
)

var dictionary = Article(
	H1("Dictionary"),
	P(`Short list of words/terms often used in software engineering
	   and sometimes defined differently. Only domain agnostic terms
	   have been listen, for the rest consult an english dictionary.
	   I often use the <code>dict</code> command line tool.`),

	Dl(
		Dt("Argument"),
		Dd("String following the command on the command line."),

		Dt("Flag"),
		Dd("Boolean option."),

		Dt("Option"),
		Dd("Argument starting with single or double dashes."),
	),
)
