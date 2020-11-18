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
			areas. Each area supports people differently based on
			their role in the space endeavour. A set of scenes
			describe various use cases to highlight and elicitate
			requirements of the final system.`),

			Section(
				H3("Simulation"),

				P(`Simulations enable people to build experience in
				handling stressful situations.`),

				Scene(`An incident occurs during reentry. The
				controller talks to astronaut working out the
				situation.`),

				Feature("Radio communication"),
				Ul(
					Li("Voice/audio is prefered way of communicating."),
					Li("Texting alternative when there is radio interference"),
				),
				//
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

		// problems concerning the purpose
		//			Problem(`Stear precisely using thrusters.`),
		//			Problem(`Adapt course for unforseen circumstance.`),*/
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

func Role(role string) *Element {
	return Span(Class("role"), role)
}

func Feature(c ...interface{}) *Element {
	return Span(Class("feature")).With(c...)
}

func Scene(c ...interface{}) *Element {
	return P(Class("scene")).With(c...)
}

func Requirement(c ...interface{}) *Element {
	return Span(Class("requirement")).With(c...)
}
