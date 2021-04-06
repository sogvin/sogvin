package sogvin

import (
	"path/filepath"

	_ "embed"

	"github.com/gregoryv/draw/design"
	"github.com/gregoryv/draw/shape"
	"github.com/gregoryv/navstar"
	"github.com/gregoryv/navstar/htapi"
	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/toc"
)

var roleBasedService = func() *Element {
	nav := Nav()
	article := Article(
		//

		H1("System design - layers and access roles"),

		P(`Services running on cloud servers mostly provide
	    programming interfaces through the hyper text transfer
	    protocol(HTTP) and are often sizeable applications. Without a
	    clear design they become hard to maintain over time. Moreover
	    their features need protecting from unauthorized access. This
	    protection can be defined in roles.  I'll explore one package
	    design and elaborate on the naming with an example service for
	    navigating the stars.`),

		nav,

		P(`Code examples found below you can view at `,
			github("gregoryv/navstar", "gregoryv/navstart"), `.`),

		H2("Domain description"),

		P(`Before we go into the design, let me tell you about the
        business of navigating through the galaxy. By describing the
        domain we'll be able to elicit concepts and features for our
        system design.`),

		Div(Class("figure"), Img(Src("img/galaxytravel.png"))),

		P(`The company <b>Future Inc.</b> provides people means to
	    travel the Milky Way. Customers, browse and order trips on `,
			Code("galaxytravel.future.now"), `. Destinations are cataloged
	    and presented adventurously with specifications of distance,
	    travel time, ship details as well as a captains profile.`,
			Br(), `The ships captain uses the same service to plan the
	    entire flight. He submits a flight plan just days before
	    departure to make sure it's as accurate as possible since
	    space travel is not an exact science and there are lot of
	    unknown objects about. Luckily the navigation system provides
	    the pilots with all the information they need. Once the plan
	    has been submitted, passengers can view route details,
	    including interesting waypoints. Crew members also access the
	    details of routes and possible alternatives, should there be
	    an unforseen cosmic event.`),

		P(`Now that we know a bit about the domain we'll be working
	    in, lets find the important concepts and focus on the ones
	    part of navigating the stars.`),

		H3(`Concept elicitation`),

		P(`We know the service is found at
	    galaxytravel.future.com. This is a domain name selected
	    because it sounds great and is easily remembered by customers
	    when they want to elope to another part of the galaxy, imagine
	    Luke Skywalker in a bar. It has very little to do with
	    navigating the stars though so we should exclude that name or
	    part of it in our design. Reason about the words in your
	    domain before deciding on how to use them. Refactoring code is
	    easy compared to changing peoples perception of concepts.`),

		P(`Several people are interacting with the service; customers,
	    captain, crew members and passengers. Let's exclude the
	    customer as that is a role more related to booking. This
	    leaves us with passenger, captain and crew members. Obviously
	    the passenger is a customer at some point but the word
	    customer is irrelevant when it comes to navigating the
	    stars. A passenger however, has access to viewing parts of the
	    flight plan, which leads us to enumerating the features of our
	    navigation service.`),

		P(`We recognize that the galaxytravel service, serves both
	    customers and captains though with different purpose. In our
	    design we'll separate these into different systems and focus
	    on the system that provides features for maninpulating
	    flightplans. The captain submits a flightplan whereas, other
	    crew members and passengers can view it. Passengers can see
	    the designated route, with details such as current location,
	    waypoints and estimated time of arrival.`),

		P(`Let's summarize by grouping the concepts`),

		Dl(
			Dt(`Role`),
			Dd(`passenger, captain, crew member`),

			Dt(`Rsource`),
			Dd(`flightplan, route, waypoint`),

			Dt(`Feature`),
			Dd(`submit flightplan, view flightplan`),
		),

		P(`Note that up until now all the terminology is from the
    	domain of navigating the stars. The only term used that
    	somehow relates to a software is "system". Which we'll design
    	now.`),

		// ----------------------------------------

		H2("System design"),

		P(`The first thing we need is a name for the package or module
	    that will contain the source code of our software.`),

		H3("Package naming"),

		P(`One way to figure out a good name is to try to write that
	    one line package documentation sentence. `, Em(`"Package X
	    provides ..."`)),

		Em(`"Package galaxytravel provides applications for planning
	    star navigation"`),

		P(`Sounds ok, but wait, we said that the service name
	    galaxytravel was selected for customers and should be excluded
	    from the navigation system. Also it as a service provides more
	    than just applications for planning star navigation. How
	    about`),

		Em(`"Package starnavigation provides means to plan galaxy flights"`),

		P(`Short sentence which abstracts what it provides by using
	    the word "means" and is more specific by using "plan galaxy
	    flights". One problem though, the name "starnavigation" is a
	    mouthful, with five syllables. The name will be used
	    extensively and we should try to find something
	    shorter. Maybe`),

		Sidenote("Short pronounce- able package name", 0.0),
		Em(`"Package navstar provides a system for planning galaxy flights"`),

		P(`Short pronouncable name, mentions the system and its main
	    purpose. It allows for easy discussion and ties into the
	    domain terminology nicely. Let's stick with it for now.`),

		// ----------------------------------------

		H3("Navstar package"),

		P(`Navstar implements domain logic related to planning galaxy
	    flights. It's at the core of our design. Later we'll build
	    other layers on top of it.`),

		Div(Class("figure"), coreDiagram(`Navstar is the core package
	    with domain logic`).Inline()),

		P(`The type system is the most prominent abstraction the
	    navstar package provides. It's responsible for synchronizing
	    database access and other domain related configuration. There
	    would usually only exist one instance of the system in any
	    running application.`),

		navrepo.LoadFile("system.go"),

		P(`Roles expose access to user methods. Fairly often we talk
	    about what we can do with a system, referring to you and me as
	    users.`),

		Em(`- Pilots submit flightplan`), Br(),
		Em(`- Passengers and crew member view flightplans`),

		P(`This translates to SubmitFlightplan is implemented by type
	    user and accessible via the pilot role. Also ListFlightplans
	    is implemented by type user but accessible by roles pilot,
	    passenger and crew member.`),

		Div(Class("figure"), navstarDiagram(`Different roles provide
		different methods`).Inline()),

		P(`We start of by defining all roles in one file together with
	    the interface, showing partial content below. The reason being
	    that roles change together, ie. if we define a new feature
	    method, all roles need updating.`),

		navrepo.LoadFile("role.go", 3, 25),

		P(`This design provides well defined places to implement
	    future features. Assume the navstar system should provide
	    planet information to users.`),

		Ol(
			Li("Define resource Planet"),
			Li("Implement feature methods on type user, e.g.",
				Ul(
					Li(Code("viewPlanet(name string)")),
					Li(Code("savePlanet(v Planet) error")),
				),
			),
			Li("Expose user methods to selected roles"),
		),

		Sidenote(`Authentication is most often a service level
		feature.`, 0.0),

		P(`Note that authentication is not part of this design,
	    i.e. translating some user credentials into one specific
	    role. The reason is that authentication is not part of the
	    navstar domain.`),

		P(`At this point the navstar system is fairly well designed
	    and we know how to extend it with new features. It's time to
	    expose the navstar system through a HTTP programming
	    interface.`),

		// ----------------------------------------

		H2("HTTP programming interface"),

		P(`At this stage I haven't decided on a specific design for
		the interface. I know however that talking about this part of
		the design; we will use wording like "navstar webapi",
		"navstar httpapi" or even simply "navstar api". Now, a name
		like webapi or api alone seems a bit to generic as HTTP is not
		the only protocol available to us. Httpapi is a mouthful so
		we'll shorten it to <em>htapi</em>.`),

		Div(Class("figure"), htapiDiagram(`htapi package is separated
		from the core navstar`).Inline()),

		P(`The htapi provides a router that exposes the navstar
	    features using its system and roles. Resources are accessible
	    via different URLs. The routing of a url to a specific server
	    method is handled by the muxer. Note how in this layer we are
	    increasingly using terms outside of the domain and more
	    technical, which is perfectly ok.`),

		navrepo.LoadFile("htapi/router.go", 0, 21),

		P(`A request from a client such as a browser would follow the
	    below sequence.`),

		usingNavstarSystem("Using navstar system via a HTTP interface").Inline(),

		P(`The router only propagates the request down to the muxer
		which is an implementation detail and can be freely replaced
		if needed. This way the router references everything needed by
		the handler functions which are bound to it.`),

		navrepo.LoadFile("htapi/router.go", 23, -1),

		P(`We can keep on developing this layer until we think it's
		ready to let other people start using it. This would be the
		time you think about designing for deployment, performance and
		maintainers. However I won't go into those areas in this
		article. Let's focus on the design for one particular
		application we intend to deliver that can be hosted on some
		server in the cloud.`),

		// ----------------------------------------

		H2("Application"),

		P(`With our domain logic in one package and the exposing http
	    layer in another we want to provide a command that handles
	    requests over the internet. The Go language build system
	    defaults to folder names when building and often these
	    command-folders are found under the folder
	    <code>cmd/</code>. We can use the same method to find a good
	    name for the package holding the application. After some
	    interations I ended up with with the name <em>starplan</em>.`),

		Div(Class("figure"), starplanDiagram(`Command starplan exposes
		the htapi via a TCP server.`).Inline()),

		P(`The reason you shouldn't name it e.g. "navstar" is that the
	    domain of navigating stars will grow and you probably want to
	    expose parts of it differently, thus having multiple
	    commands.`, Br(), `Adding files for some of the mentioned
	    abstractions we end up with a directory tree like this`),

		ShellCommand("$ tree navstar\n"+navstarTree),

		H2("Summary"),

		P(`Separating the domain logic from, the application exposing
	    it, allows your service to grow more easily. By naming
	    components carefully we can reason about concepts such as
	    the-galaxytravel-service, navstar-system and
	    starplan-application, which are all easily referencable in the
	    source code aswell. Aim to have a few, but well defined
	    crossing points between the layers. Starplan uses htapi and
	    navstar, whereas htapi only uses the navstar package.`),

		Div(Class("figure"), InternalDiagram(`Dependency flow, from
		right to left.`).Inline()),

		P(`For other internal domain logic that benefits from
		alternate naming than <code>navstar.X</code>, structure
		packages in the internal directory. A part from being hidden
		by the Go language it's also conceptually correct that domain
		internals should not be exposed to any other layer.`),

		P(`Try not to design all the layers simultaneously as it's
		easier to reason about one purpose. Start with the business
		domain logic and work outwards throught the layers.`),

		Em(`See you in the stars!`),

	//
	)
	toc.MakeTOC(nav, article, "h2", "h3")
	return article
}()

var repoLink = A(
	Href("https://github.com/gregoryv/navstar"),
	"github.com/gregoryv/navstar",
)

//go:embed "example/navstar.tree"
var navstarTree string

func coreDiagram(caption string) *design.Diagram {
	var (
		w, h, r, s = 80, 50, 20, 2
		dx         = w - r
		dy         = h / 2
		right      = dx + 2*s
		below      = -dy + s
		above      = dy - s
		d          = design.NewDiagram()
		ns         = shape.NewHexagon("navstar", w, h, r)
		htapi      = shape.NewHexagon("", w, h, r)
		cmd        = shape.NewHexagon("", w, h, r)
		starplan   = shape.NewHexagon("", w, h, r)
	)
	_ = below
	shape.SetClass("dim", htapi, cmd, starplan)
	d.Place(ns).At(80, 120)

	d.Place(htapi).Above(ns, 0)
	shape.Move(htapi, right, above)

	d.Place(cmd).Above(ns, 2*s)
	d.Place(starplan).Above(cmd, 0)
	shape.Move(starplan, right, above)

	d.SetCaption(caption)
	return d
}

func htapiDiagram(caption string) *design.Diagram {
	var (
		w, h, r, s = 80, 50, 20, 2
		dx         = w - r
		dy         = h / 2
		right      = dx + 2*s
		below      = -dy + s
		above      = dy - s
		d          = design.NewDiagram()
		ns         = shape.NewHexagon("navstar", w, h, r)
		htapi      = shape.NewHexagon("htapi", w, h, r)
		cmd        = shape.NewHexagon("", w, h, r)
		starplan   = shape.NewHexagon("", w, h, r)
	)
	_ = below
	shape.SetClass("dim", ns, cmd, starplan)
	d.Place(ns).At(80, 120)

	d.Place(htapi).Above(ns, 0)
	shape.Move(htapi, right, above)

	d.Place(cmd).Above(ns, 2*s)
	d.Place(starplan).Above(cmd, 0)
	shape.Move(starplan, right, above)

	d.SetCaption(caption)
	return d
}

func starplanDiagram(caption string) *design.Diagram {
	var (
		w, h, r, s = 80, 50, 20, 2
		dx         = w - r
		dy         = h / 2
		right      = dx + 2*s
		below      = -dy + s
		above      = dy - s
		d          = design.NewDiagram()
		ns         = shape.NewHexagon("navstar", w, h, r)
		htapi      = shape.NewHexagon("htapi", w, h, r)
		cmd        = shape.NewHexagon("cmd", w, h, r)
		starplan   = shape.NewHexagon("starplan", w, h, r)
	)
	_ = below
	shape.SetClass("dim", ns, htapi, cmd)
	d.Place(ns).At(80, 120)

	d.Place(htapi).Above(ns, 0)
	shape.Move(htapi, right, above)

	d.Place(cmd).Above(ns, 2*s)
	d.Place(starplan).Above(cmd, 0)
	shape.Move(starplan, right, above)

	d.SetCaption(caption)
	return d
}

func InternalDiagram(caption string) *design.Diagram {
	var (
		w, h, r, s = 80, 50, 20, 2
		dx         = w - r
		dy         = h / 2
		right      = dx + 2*s
		left       = -right
		below      = -dy + s
		above      = dy - s
		d          = design.NewDiagram()
		ns         = shape.NewHexagon("navstar", w, h, r)
		htapi      = shape.NewHexagon("htapi", w, h, r)
		cmd        = shape.NewHexagon("cmd", w, h, r)
		starplan   = shape.NewHexagon("starplan", w, h, r)
		internal   = shape.NewHexagon("internal", w, h, r)
		other      = shape.NewHexagon("other", w, h, r)
		arrow      = shape.NewArrow(200, 0, 45, 200)
	)

	shape.SetClass("dim", ns, htapi, cmd, starplan)
	d.Place(ns).At(80, 120)

	d.Place(htapi).Above(ns, 0)
	shape.Move(htapi, right, above)

	d.Place(cmd).Above(ns, 2*s)
	d.Place(starplan).Above(cmd, 0)
	shape.Move(starplan, right, above)

	d.Place(internal).Below(ns, 0)
	shape.Move(internal, left, below)

	d.Place(other).Below(internal, 2*s)
	d.Place(arrow).At(320, 45)
	d.SetCaption(caption)
	return d
}

func navstarDiagram(caption string) *design.ClassDiagram {
	var (
		d         = design.NewClassDiagram()
		role      = d.Interface((*navstar.Role)(nil))
		pilot     = d.Struct(navstar.Pilot{})
		passenger = d.Struct(navstar.Passenger{})
		crew      = d.Struct(navstar.Crew{})
	)
	d.Place(role).At(100, 20)
	d.Place(pilot).Below(role, 70)
	shape.Move(pilot, -100, 0)
	d.Place(passenger, crew).RightOf(pilot, 70)
	d.VAlignCenter(passenger, role)

	d.SetCaption(caption)
	return d
}

func usingNavstarSystem(caption string) *design.SequenceDiagram {
	var (
		d       = design.NewSequenceDiagram()
		browser = d.Add("browser")
		router  = d.AddStruct(htapi.Router{})
		role    = d.AddInterface((*navstar.Role)(nil))
		user    = d.AddStruct(navstar.User{})
		sys     = d.AddStruct(navstar.System{})
	)
	d.ColWidth = 140
	d.Link(browser, router, "GET /flightplans")
	d.Link(router, router, "serveFlightplans() via muxer")
	d.Link(router, role, "new: role")
	d.Link(router, user, "new: user")
	d.Link(router, user, "user.Use(system, role)")
	d.Link(router, role, "ListFlightplans()")
	d.Link(role, user, "listFlightplans()")
	d.Link(user, sys, "query database")
	d.Link(router, browser, "write http response")

	d.Group(router, role, "Protected by role", "blue")
	d.Group(role, sys, "Unprotected", "red")
	d.SetCaption(caption)
	return d
}

func navstarDir(subpath string) string {
	return filepath.Join("..", "navstar", subpath)
}

// ----------------------------------------
var navrepo = &Repo{
	host:  "https://github.com/gregoryv/navstar",
	local: "/home/gregory/src/github.com/gregoryv/navstar",
}
