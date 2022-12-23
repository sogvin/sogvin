package website

import (
	_ "embed"

	. "github.com/gregoryv/web"
)

func projectLayout() *Element {
	return Article(
		H1("Golang project layout"),

		P(`In this context `, Em("layout"), ` is about placing files
		and directories in a repository. `,
			A(Href("https://go.dev/blog/organizing-go-code"), `Organizing
		Go Code`), `, by Andrew Gerard, 2012 is still very much valid
		and we build on it's description. I use the term layout over
		structure becuse structure implies rigidity and software
		development needs to be decoupled and agile even at this
		level. Projects evolve and your layout should support this
		evolution. The goal is to allow for refactoring with minimal
		impact.`),

		sidenote(`Layout should minimize refacto- ring impact.`, -1.8),

		P(`Here is an example of a project layout from `,
			linkTo(roleBasedService()), "."),

		shellCommand(`$ tree ../navstar/
├── cmd
│   └── starplan
│       ├── main.go
│       └── main_test.go
├── go.mod
├── go.sum
├── htapi
│   ├── router.go
│   └── router_test.go
├── LICENSE
├── package.go
├── README.md
├── resource.go
├── role.go
├── system.go
├── system_test.go
├── user.go
└── user_test.go`),

		P(`This layout was however not created from start, it
		evolved. Initial layout was simply`),

		shellCommand(`$ tree ../navstar/
├── go.mod
├── LICENSE
└── README.md`),

		P(`a resonable base of any repository, and works well
		on github.com which displays the README.md file. The LICENSE
		is important once you publish your work and you want it to be
		accessible through `, A(Href("https://pkg.go.dev"),
			"pkg.go.dev"), `. Writing domain logic in the root of your
		repository allows for`),

		Ul(
			Li("quick access to files"),
			Li("easy refactoring with tools like gofmt"),
		),

		P(`If your repository only provides domain logic this could be
		all you need, a flat and easy to work with layout. Our example
		evolved into `),

		shellCommand(`$ tree ../navstar/
├── go.mod
├── go.sum
├── LICENSE
├── package.go
├── README.md
├── resource.go
├── role.go
├── system.go
├── system_test.go
├── user.go
└── user_test.go`),

		P(`At some point you it's time to use the domain logic. If you
		look carefully we already have couple of test files. They are
		the initial use and should contain some examples of it's
		use. Usually online systems involve some kind of remote API
		and in our example it was placed in directory named `,
			Code("htapi"), `.`),

		shellCommand(`$ tree ../navstar/
├── go.mod
├── go.sum
├── htapi
│   ├── router.go
│   └── router_test.go
├── LICENSE
├── package.go
├── README.md
├── resource.go
├── role.go
├── system.go
├── system_test.go
├── user.go
└── user_test.go`),
		//

		P(`Keep it simple and readable. Readable here means more
		referencable when discussing about the solution. E.g.`),

		Em(`- navstar has a public HTTP API`),

		P(`Let's discuss this, because an argument could be made that
		the layout should have been something like`),

		Pre(`navstar/
  http/
    api/
      router.go
`),

		P(`There are several reasons why this makes evolution more
		cumbersome. FIRST; we've added an extra, unused level with the
		http directory. It makes it slightly harder to work with
		depending on which tools you use. In the terminal it adds to
		the navigation between directories. SECOND; the name
		<em>api</em> is too generic and should we decide later to add
		another protocol, e.g. gRPC we'd have to (a) rename api or (b)
		move it to something like navstart/http/api/rest, increasing
		levels even more. THIRD; the name http is already used in the
		standard package net/http. Avoid reusing package names as
		renaming them later is harder, unless you have really
		sophisticated refactoring tools.`),

		P(`Once it's time to build an application for others to use,
		it's benefitial to place this in a directory named the same as
		the Go build tools default to naming the binary to it. Create
		your application directory under cmd/, like this`),

		shellCommand(`$ tree ../navstar/
├── cmd
│   └── starplan
└── htapi`),

		P(`Now this goes against the second argument we just made to
		exclude unused level of directories. And you are right however
		other reasons come into play here. FIRST; when sharing
		multiple commands using the Go tooling we can simplify
		installation instructions, e.g. `),

		shellCommand(`$ go install github.com/sogvin/navstar/cmd/...@latest`),

		P(`SECOND; this layout works in a conversation`),

		Em(`- I can access the navstar system through the command
		starplan`),

		P(`THIRD; this layout is followed by the Go project so it's
		familiar. FOURTH; if kept separate from the domain logic,
		commands can easily be moved to their own repositories,
		e.g. for having a separate release history.`),
		//

		H2("Internal"),

		P(`The internal directory is exactly what it says,
		internal. It is internal to a whole repository, and if treated
		like that you may end up in issues when refactoring down the
		road. If the internal directory is only imported by it's
		parent then you are all good but if other subdirectories
		import it those subdirectories cannot be moved as the cross
		package import of internal directories is not allowed unless
		within the same repository.`),

		P(`Example; navstart/internal can be imported by anything
		navstart/..., however say navstar/cmd/starplan imports it,
		then we can no longer move cmd/startplan to it's own
		repository as that import will be disallowed. Think of
		internal as internal to the parent directory only. Also use it
		only when really needed. It's easier to hide constructs such
		as types and functions withing the same package first. It
		minimizes the level management as we've mentioned earlier.`),

		H2("Summary"),

		P(`Project layouts evolve; keep a flat layout and add
		directories when needed, thinking in layers. Be mindful of
		readability and future refactorings. Keep commands in cmd/
		subdirectories and postpone use of internal by hiding
		constructs withing a package first.`),
		//
	)
}

func purposeOfFuncMain() *Element {
	return Article(
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

		loadFile("./internal/cmd/countstars/main.go", 8, -1),

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

		sidenote("Cyclomatic complexity should be one.", -5.2),
		sidenote("Option order should match output.", -2.3),

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

		sidenote("Run is now testable and complexity can grow slightly", 8.5),
		loadFile("./example/cmd/starcounter/starcounter.go"),

		P(`Complexity of func main has grown slightly, looking like`),
		sidenote(
			Span("Alternate ",
				A(
					Href("https://godoc.org/github.com/gregoryv/cmdline"), "cmdline"),
				" package for parsing arguments.",
			),
			1.8,
		),
		loadFile("./example/cmd/starcounter/main.go"),

		P(`Testing complex patterns is still doable by testing the
		main func. Although with this design unit tests of `, Code(`func
		StarCounter.Run()`), ` are probably even simpler once variation
		increase.`),

		loadFile("./example/cmd/starcounter/starcounter_test.go"),
	)
}

func nexusPattern() *Element {
	return Article(
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

		sidenote("The err field links operations.", 0.6),
		sidenote("Each method sets x.err before returning.", 3.3),
		loadFile("./internal/errhandling/nexus.go", 21, -1),

		`With the fileIO nexus inplace the CopyFile function is
	readable and with only one error checking and handling needed.`,
		loadFile("./internal/errhandling/nexus.go", 8, 19),
	)
}

func gracefulServerShutdown() *Element {
	return Article(
		H1("Graceful server shutdown"),
		P(`Avoid disrupting ongoing requests by shutting down
	   gracefully. In the below example Ctrl-c can be used to signal
	   an interrupt which tells a listening <code>http.Server</code>
	   to shutdown.`),

		sidenote("Register the graceful part of the server.", 4.8),
		sidenote("Important to wait for graceful stop to end.", 7.8),
		loadFile("./internal/cmd/graceful/graceful.go", 11, -1),
		P(`Remember that you could expose the Shutdown func of your
       server through an URL to simplify clean shutdown. Useful for
       when you are doing continuous integration and
       deployment.`),
	)
}

func componentsDiagram() *Element {
	return Article(
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
}

func strictMode() *Element {
	return Article(
		H1("Strict mode"),

		P(`Failing early is good for many reasons and strict mode design
       makes this very helpful during testing. An http client is a
       good place for this design.`),

		H2("Client"),

		P(`HTTP clients issue requests to some service, mainly by the Do
       method. Define the Strict interface to match that of the
       familiar testing T.Fatal.`),

		loadFile("./internal/strictClient.go", 8, 20),

		P(`Once the client has the strict ability it can be used in it's
	  methods. Default the client to a lax mode where the Fatal method
	  does nothing.  `),

		loadFile("./internal/strictClient.go", 22, 28),

		P(`Let's assume your service only accepts json and expects each
       request to set the correct header.  A simple wrapper around
       http.DefaultClient could look like this.  `),

		sidenote("Use the strict wrapper in public methods.", 1.2),
		sidenote("Private funcs just return errors as usual.", 5.7),
		loadFile("./internal/strictClient.go", 30, -1),

		P(`Any error from the sending of the request will be checked by
	  the strict interface. This adds no real benefit to the client
	  itself but it makes a difference when testing.`),

		loadFile("./internal/strictClient_xtest.go", 8, 13),

		shellCommand(`$ go test
--- FAIL: TestClient (0.00s)
    strictClient.go:32: checkContentType: "" must be application/json
`),

		P(`Descriptive error messages make tests short and concise.  Use
	   <em>check</em> prefix to distinguish from asserts.`),
	)
}
