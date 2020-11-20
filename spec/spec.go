package spec

import (
	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/toc"
)

func NewSpecification() *Page {

	nav := Nav()
	spec := Article(
		H1("Spaceship navigation system"),
		Em(`Purpose; provide safe travel through space.`),

		P(`The navigation system provides people a way to plot a
           course through space or manually steer a ship.  People
           depend on its accuracy and automation to safely navigate
           through space.`),

		nav,
		Section(
			H2("Domain"),

			P(`This specification is divided into domains. Each
			supports people differently based on their role in the
			space endeavour. A set of scenes describe various use
			cases to highlight and elicitate requirements of the final
			system.`),

			H3("Navigation"),

			H4(Scene(`Plot new course`)),

			P(`Standing at the bridge, the captain asks for the
			closest viable planets for some time at the beach. Selects
			the one with the nicest beaches and views and tells the
			system to plot the course. The plot details show there
			route goes through uncharted space. The captain selects
			another of the viable planets and tells the system to plot
			the course again. Once satisfied, he tells the system to
			engage.`),

			P(`The journey is estimated to five days. On the second
			day however an interference is detected in space and the
			ship adapts the course accordingly. The captain is
			notified through his personal communicator of the
			changes.`),
			//

			H3("Control"),

			P(`Manual override`),
		),
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
