package sogvin

import (
	_ "embed"
	"net/http"

	"github.com/gregoryv/sogvin/example/spaceflight"
	"github.com/gregoryv/sogvin/example/spaceflight/cmd/htspace"

	"github.com/gregoryv/draw/design"
	"github.com/gregoryv/draw/shape"
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

var roleBasedService = Article(
	H1("Role based service"),

	P(`Servers provide a service through some protocol, often the
	protocol is HTTP. Services like this are often sizeable
	applications and without a design they become hard to maintain
	over time. Moreover services and their features need protecting
	from unauthorized access. One approach is to provide role based
	access.`),

	Sidenote(`Use different nam- es for package, commands and DNS.`,
		2.0),

	P(`Let's assume we are writing a service for spaceflight. This
    service will be hosted on `, Code("galaxytravel.future.now"),
		`. The domain logic is `, Em("spaceflight"), ` i.e. the package
    name. We'll name the application providing this service via HTTP, `,
		Em("htspace"), ".", Br(), `Don't name the application the same as
    the domain logic package. Also refrain from naming it same as the
    DNS name where it's hosted. The DNS will remain for a long time
    whereas your system will evolve, be split up into smaller
    applications with specific responsibilities. The domain logic
    however will probably remain the same. The directory layout looks
    like this`),

	ShellCommand("$ tree spaceflight\n"+spaceflightTree),
	//

	P(`The system is the most prominent abstraction the spaceflight
	package provides. It's responsible for synchronizing database
	access and other domain related configuration. There would usually
	only exist one instance of the system.`),

	LoadFullFile("", "./example/spaceflight/system.go"),

	P(`Roles expose access to user methods. Fairly often we talk about
	what we can do with a system, referring to you and me as
	users.`),

	Em(`- Pilots plan routes`), Br(),
	Em(`- Passengers and pilots view routes`),

	P(`This translates to PlanRoute is implemented by type user and
	accessible via the pilot role. Also ListRoutes is implemented by
	type user but accessible by both roles pilot and passenger.`),

	Div(Class("figure"), spaceflightDiagram(`Different roles provide
		different methods`).Inline()),

	LoadFullFile("", "./example/spaceflight/role.go"),

	P(`This design provides well defined places to implement future
	features. Assume the spaceflight service should provide planet
	information to users.`),

	Ol(
		Li("Define resource Planet"),
		Li("Implement read write methods on type user, e.g.",
			Ul(
				Li(Code("viewPlanet(name string)")),
				Li(Code("savePlanet(v Planet) error")),
			),
		),
		Li("Expose user methods to selected roles"),
	),

	P(`Once you need authentication you have the option to make it
	part of this service, by extending the Service.Use() method with
	e.g. credentials argument.`),

	H2("HTTP interface"),

	P(`The application htspace can now expose the spaceflight features
	using its system and roles. An application provides methods for
	accessing resources via different URLs. The routing of a url to a
	specific server method is handled by the subsequent router.`),

	LoadFullFile("", "./example/spaceflight/cmd/htspace/application.go"),

	P(`A request from a client such as a browser would follow the
	below sequence.`),

	usingSpaceflightSystem("Using spaceflight system via a HTTP interface").Inline(),

	P(`Separating the domain logic from the application exposing it
	using some protocol allows your service to grow. Naming components
	carefully we can reason about concepts such as the-galaxytravel-service,
	spaceflight-system and htspace-application, which are all
	easily referencable in the source code aswell.`),
)

//go:embed "example/spaceflight.tree"
var spaceflightTree string

func spaceflightDiagram(caption string) *design.ClassDiagram {
	var (
		d         = design.NewClassDiagram()
		role      = d.Interface((*spaceflight.Role)(nil))
		pilot     = d.Struct(spaceflight.Pilot{})
		passenger = d.Struct(spaceflight.Passenger{})
	)
	d.Place(role).At(120, 20)
	d.Place(pilot).Below(role, 70)
	shape.Move(pilot, -100, 0)
	d.Place(passenger).RightOf(pilot, 70)

	d.SetCaption(caption)
	return d
}

func usingSpaceflightSystem(caption string) *design.SequenceDiagram {
	var (
		d       = design.NewSequenceDiagram()
		browser = d.Add("browser")
		srv     = d.AddStruct(http.Server{})
		app     = d.AddStruct(htspace.Application{})
		role    = d.AddInterface((*spaceflight.Role)(nil))
		sys     = d.AddStruct(spaceflight.System{})
	)
	d.ColWidth = 140
	d.Link(browser, srv, "GET /routes")
	d.Link(srv, app, "serveRoutes()")
	d.Link(app, role, "new: role")
	d.Link(app, sys, "role.ListRoutes()")
	d.Link(app, browser, "write http response")

	d.Group(app, sys, "Role based access to domain logic", "blue")
	d.SetCaption(caption)
	return d
}

// ----------------------------------------
