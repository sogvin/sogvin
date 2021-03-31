package sogvin

import (
	_ "embed"

	. "github.com/gregoryv/web"
)

var purposeOfFuncMain = Article(
	H1("Purpose of <code>func main()</code>"),

	P(`The purpose of <code>func main()</code> is to <b>translate
       commandline arguments to application startup state</b>. Once
       the state is prepared a specific entry function is called. More
       often than not, logging verbosity is one such state that needs
       to be configured early on.  <br> Go provides the builtin flag
       package to define, document and parse the arguments.`),

	H2("Example <code>CountStars(galaxy string)</code>"),

	P(`Imagine an application that counts the stars in a named
	   galaxy. The main function should then make sure the options are
	   correct and forward them as arguments to the function doing the
	   actual work. The name of the galaxy would be such an option and
	   perhaps a verbosity flag for debugging purposes.`),

	LoadFile("./internal/cmd/countstars/main.go", 8, -1),

	P(`Now that you know what the main function should do, let us take
	   a look at how to do it, apart of the option definition
	   and argument passing.<br> First, the cyclomatic complexity of
	   the main function is one. Ie. there is only one path through
	   this program.  There are however two exit points, apart from
	   the obvious one <code>flag.Parse()</code> exits if the parsed
	   options do not match the predefined. The single pathway means
	   that testing the main function is simple. Execute this
	   application with valid arguments and all lines are covered, leaving
	   all other code for unittesting.<br> Also, if you execute the
	   program you would note that second, the order of the options are
	   sorted in the same way as the help output.`),

	Sidenote("Cyclomatic complexity should be one.", -5.2),
	Sidenote("Option order should match output.", -2.3),

	H2("Benefits"),

	P(`Adhering to the &ldquo;keep it simple principle&rdquo; and only
	   doing one thing in each function, works out nicely for the main
	   function as well. One could argue that, if you moved everything
	   inside main into a start function, the option definitions would
	   also be tested.  Think about it for a minute and figure out
	   what exactly you would be testing. If the flag package already
	   makes sure it's functions work as expected the only thing left
	   is testing which options you have defined.  They would need to be
	   updated each time you add or remove an option  which is a sign of a
	   poor test.<br> You could potentially refactor main and separate
	   the option definitions into smaller functions for readability but
	   you still wouldn't need to write unittests for them. However,
	   when your application grows and command line arguments start
	   having relations you ought to verify that. More on this in the
	   next section.`),

	P(`But start of and keep main simple, constrain it to only set
	   global startup state before calling the one function that does
	   the actual work.<br>This works great for services and simpler
	   commands that only do one thing.`),

	H2("More advanced commands"),

	P(`When the commands get more complex with many more options the
	   above approach has its limits. Number of arguments to
	   CountStars will grow and become hard to verify any relations
	   between them. One solution is to turn func CountStars into a
	   command. Advanced commands may also have logic for combination
	   of options which would suggest you should verify command
	   execution with various options. This is impossible to do with
	   the above approach while tracking coverage.`),

	Sidenote("Run is now testable and complexity can grow slightly", 5),
	Sidenote(
		Span("Alternate ",
			A(
				Href("https://godoc.org/github.com/gregoryv/cmdline"), "cmdline"),
			" package for parsing arguments.",
		),
		13,
	),

	LoadFile("./example/cmd/starcounter/starcounter.go"),

	P(`Testing complex patterns is straight forward.`),
	LoadFile("./example/cmd/starcounter/starcounter_test.go"),

	P(`Complexity of func main still remains at one, looking like`),
	LoadFile("./example/cmd/starcounter/main.go"),
)

var nexusPattern = Article(
	H1("Nexus pattern"),
	P(
		"The word nexus is defined as ",
		Quote(
			"&#34;The means of connection between things linked in series&#34;",
		),
		". The pattern is useful in ",
		A(
			Href(
				"https://go.googlesource.com/proposal/+/master/design"+
					"/go2draft-error-handling-overview.md",
			),
			"error handling",
		),
		" sequential function calls.",
	),

	H2("Example <code>CopyFile(from, to string)</code>"),
	P(`Copying a file, if done all in one function, is unreadable due
       to multiple error checking and handling.  With the nexus
       pattern you define a <code>type fileIO struct</code> with the
       error field. Each method must check the previous error and
       return if it is set without doing anything. This way all
       subsequent calls are no-operations.`),

	Sidenote("The err field links operations.", 0.6),
	Sidenote("Each method sets x.err before returning.", 3.3),
	LoadFile("./internal/errhandling/nexus.go", 21, -1),

	`With the fileIO nexus inplace the CopyFile function is
	readable and with only one error checking and handling needed.`,
	LoadFile("./internal/errhandling/nexus.go", 8, 19),
)

var gracefulServerShutdown = Article(
	H1("Graceful server shutdown"),
	P(`Avoid disrupting ongoing requests by shutting down
	   gracefully. In the below example Ctrl-c can be used to signal
	   an interrupt which tells a listening <code>http.Server</code>
	   to shutdown.`),

	Sidenote("Register the graceful part of the server.", 4.8),
	Sidenote("Important to wait for graceful stop to end.", 7.8),
	LoadFile("./internal/cmd/graceful/graceful.go", 11, -1),
	P(`Remember that you could expose the Shutdown func of your
       server through an URL to simplify clean shutdown. Useful for
       when you are doing continuous integration and
       deployment.`),
)

var componentsDiagram = Article(
	H1("Components diagram"),
	P(
		`The components diagram shows services and processes related
		to one another. Good for system overviews and microservice
		architectures.`,
	),
	Span(Class("Center"),
		newOverviewDiagram(),
	),
	P(
		`Use lines between components unless you are conveying
		signaling direction between them.`,
	),
	H2("General diagram tips"),
	P(
		`In general diagrams should be kept simple, 5-9 items is a
		good rule (same as complexity in code).`,
	),
	Ul(
		Li("Highlight important components"),
		Li("Use arrows when direction is important otherwise plain lines"),
		Li(
			`Reflect on cognitive placement, ie. cloud components are
			above others`,
		),
		Li(
			`Use white as emphasizing color, works in both grayscale
			 and colored diagrams`,
		),
		Li("Stick to one color scheme"),
	),
	colorSchemeDiagram(),
)

var strictMode = Article(
	H1("Strict mode"),

	P(`Failing early is good for many reasons and strict mode design
       makes this very helpful during testing. An http client is a
       good place for this design.`),

	H2("Client"),

	P(`HTTP clients issue requests to some service, mainly by the Do
       method. Define the Strict interface to match that of the
       familiar testing T.Fatal.`),

	LoadFile("./internal/strictClient.go", 8, 20),

	P(`Once the client has the strict ability it can be used in it's
	  methods. Default the client to a lax mode where the Fatal method
	  does nothing.  `),

	LoadFile("./internal/strictClient.go", 22, 28),

	P(`Let's assume your service only accepts json and expects each
       request to set the correct header.  A simple wrapper around
       http.DefaultClient could look like this.  `),

	Sidenote("Use the strict wrapper in public methods.", 1.2),
	Sidenote("Private funcs just return errors as usual.", 5.7),
	LoadFile("./internal/strictClient.go", 30, -1),

	P(`Any error from the sending of the request will be checked by
	  the strict interface. This adds no real benefit to the client
	  itself but it makes a difference when testing.`),

	LoadFile("./internal/strictClient_xtest.go", 8, 13),

	ShellCommand(`$ go test
--- FAIL: TestClient (0.00s)
    strictClient.go:32: checkContentType: "" must be application/json
`),

	P(`Descriptive error messages make tests short and concise.  Use
	   <em>check</em> prefix to distinguish from asserts.`),
)

// ----------------------------------------
