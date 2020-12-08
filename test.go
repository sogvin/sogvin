package sogvin

import . "github.com/gregoryv/web"

var inlineTestHelpers = Article(
	H1("Inline test helpers"),
	P(

		`Use inline test helpers to minimize indentation and have
         failures point out failed cases directly. Given a function
         calculating the double of an int.`,
	),
	LoadFile("./internal/testing/inline/double.go", 7, 0),
	P(

		`The test would look like this.`,
	),
	Sidenote("Inlined helper does not need t argument.", 0.8),
	Sidenote("Descriptive cases fail on correct line.", 4.6),
	LoadFile("./internal/testing/inline/double_test.go", 7, -1),
	Sidenote("Utmost 2 inlined helpers.", 0.2),

	P(`Keep it simple and use utmost two inlined helpers. Compared to
       table-driven-tests inlined helpers declare the <em>how</em>
       before the cases.  If you have many cases, this style is more
       readable as you first tell the reader the meaning of
       &#34;ok&#34; and &#34;bad&#34;.  <br> Another positive benefit
       of this style is values are not grouped in a testcase
       variable. I.e. readability improves as the values are used
       directly.  <br>This style may be less readable if each case
       requires many values, though it depends on the lenght of the
       values combined.`),
)

var alternateDesign = Article(
	H1("Alternate design to simplify tests"),
	P(
		`Testing existing code you have several options to write sleek
		tests. Table driven or inlined test helpers work nicely. When
		writing new code however you have the option to choose a
		design that will be easier to verify. One go idiom is to
		return a value with an error. What if you didn't follow that
		idiom?`,
	),
	Ul(
		Li("what if you always used panics?"),
		Li("what if you only returned a struct with an optional error field?"),
	),
	P(

		`Don't let the idiom stop you from experimenting. While
         working with inline helpers I found that functions, which
         only return errors, resulted in simpler and more readable
         tests. Two assert functions are needed, one for checking for
         an error and the other for nil errors. Remember that tests
         should focus on verifying logic, not data. In this case the
         logic is binary, failed or not.`,
	),
	LoadFile("./internal/testing/okbad/assert_test.go", 8, 0),
	P(

		`The initial design of the `, A(
			Href("inline_test_helpers.html"),
			"function double",
		),

		` follows the go idiom of returning a value with an error.
         Redesign the function to take the resulting argument and only
         return an error adds a few more lines to the function. We
         also added the check for nil result. The nil check may be
         left out or removed once you have your tests.  `),
	LoadFile("./internal/testing/okbad/double.go", 7, 24), P(

		`Let's use our new assert functions.`,
	),
	LoadFile("./internal/testing/okbad/double_test.go", 7, 0),
)
