package spec

import (
	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/toc"
)

func NewSpecification() *Page {

	nav := Nav()
	spec := Article(
		H1("Spaceship control system"),
		Em(`Purpose; provide safe navigation through space.`),

		P(`The control system of a spaceship is at the heart of the
		ship. Everything on the ship is somehow monitored or
		controlled by the system. People depend on its accuracy and
		automation to safely navigate through space.`),

		nav,
		Section(
			H2("Areas"),

			P(`This specification is divided into four main
			areas. Each area supports people in different ways
			depending on their role in the space flight.  Astronauts
			need to practice their skills with the spaceship and
			controllers different aspects of the monitoring
			system. Each area is further describe with possible scenes
			to highlight and elicitate requirements of the final
			system.`),

			Section(
				H3("Simulation"),
			),
			Section(
				H3("Automation"),
			),
			Section(
				H3("Remote control"),
			),
			Section(
				H3("Maintenance"),
			),
		),

		/*

			// roles interacting with the system
			Role("Astronaut"),
			Role("Controller"),
			Role("Engineer"),

			// problems concerning the purpose
			Problem(`Stear precisely using thrusters.`),
			Problem(`Adapt course for unforseen circumstance.`),*/
	)
	//CheckRoles(spec, t.Error)

	toc.MakeTOC(nav, spec, "h2", "h3")

	page := NewPage(
		Html(
			Head(
				Style(Theme()),
			),
			Body(
				spec,
			),
		),
	)
	return page
}
