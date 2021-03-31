package sogvin

import (
	"net/http"
	"path/filepath"

	_ "embed"

	"github.com/gregoryv/draw/design"
	"github.com/gregoryv/draw/shape"
	"github.com/gregoryv/spaceflight"
	"github.com/gregoryv/spaceflight/cmd/htspace"
	. "github.com/gregoryv/web"
)

var roleBasedService = Article(
	H1("Role based service"), // todo find a better name

	P(`Servers provide a service through some protocol, often the
	protocol is HTTP. Services like this are often sizeable
	applications and without a design they become hard to maintain
	over time. Moreover services and their features need protecting
	from unauthorized access. One approach is to provide role based
	access design.`),

	H2("Navigating the stars"),

	P(`Before we go into the design, let me tell you a bit about the
    business of navigating through the galaxy. By describing the
    domain we'll be able to elicit concepts and features for our
    system design.`),

	P(`The company <b>Future Inc.</b> provides people means to travel
    the Milky Way. Customers, browse and order trips on `,
		Code("galaxytravel.future.now"), `. Destinations are cataloged and
    presented adventurously with specifications of distance, travel
    time, ship details as well as a captains profile.`, Br(),

		`Captains submit a flight plan just days before departure to
	make sure it's as accurate as possible since space travel is not
	an exact science and there are lot of unknown objects
	about. Luckily the navigation system provides them with all the
	information they need. Once the plan has been submitted,
	passengers can view route details, including interesting
	waypoints. Crew members can also access the details of routes and
	possible alternatives, should there be an unforseen cosmic
	event.`),

	H3(`Elicitation`),

	P(`We have enough information to get started I think. The main
	domain concepts are`),

	Ul(
		Li("navigating the stars - the domain"),
		Li("galaxytravel.future.com - service"),
	),

	P(`Focusing on the navigation part there are people(users) using
	the system with the following roles`),

	Ul(
		Li("captain"),
		Li("passenger"),
		Li("crew member"),
	),

	P(`The system provides features to manipulate resources`),

	Ul(
		Li("catalog, destination"),
		Li("ship"),
		Li("flight plan, route, waypoint"),
		Li("system"),
	),

	H2("Package design"),

	P(`The domain we are working in is navigating the stars, as we
	found out during our elicitation. Let's abbreviate it in a short,
	pronouncable name `, Em("navstar"), `. This will be the package
	name in our design.`),

	P(`The service is hosted on galaxytravel.future.com. This is a
	commercial name for a booking system and in our case also the
	interface to the navigation system. As such, decouple the
	commercial name from anything in your design, i.e. don't use
	it.`),

	P(`The application, that will expose the navstar features via
	HTTP, we'll name `, Em("htspace"), ".", ` The reason you shouldn't
	name it e.g. "navstar" is that the domain of navigating stars will
	grow and you probably want to expose parts of it differently, thus
	having multiple applications. Also "galaxytravel" is a poor name
	as it's the commercial service name which could be composed of
	many applications. Also the DNS will remain for a long time
	whereas your system will evolve, be split up into smaller
	applications with specific responsibilities. Adding files for some
	of the mentioned abstractions we end up with a directory tree like
	this`),

	ShellCommand("$ tree spaceflight\n"+spaceflightTree),
	//

	P(`The system is the most prominent abstraction the spaceflight
	package provides. It's responsible for synchronizing database
	access and other domain related configuration. There would usually
	only exist one instance of the system.`),

	LoadFullFile("", navstarDir("system.go")),

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

	LoadFullFile("", navstarDir("role.go")),

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

	LoadFullFile("", navstarDir("cmd/htspace/application.go")),

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

func navstarDir(subpath string) string {
	return filepath.Join("..", "spaceflight", subpath)
}
