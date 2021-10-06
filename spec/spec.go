package spec

import (
	. "github.com/gregoryv/web"
)

func NewSpecification() *Page {
	return newPage(
		H1("Exploring requirements engineering"),

		Article(
			NewExploreRequirementsEngineering(),
			NewBeachStory(),
		),
	)
}

func NewExploreRequirementsEngineering() *Element {
	return Wrap(

		P(`An exercise in elicitating requirements, imho. still one of
        the most difficult task in software engineering.`),

		P(`As a software engineer you are tasked to produce software
		systems to fulfill the need of a stakeholder. I use the term
		software engineer, or just engineer, for all roles used today
		in the industry that somehow contribute to producing
		software. The reason is they all have one thing incommon, they
		have to understand the purpose of their work. Without it, the
		end result will never be as good as envisioned by the
		stakeholder.`),

		P(`As an engineer I solve problems. One reoccuring problem is
		the difficulty of conveying knowledge from stakeholders to the
		engineer, in such a manner that it is easily understood.
		There are many reasons for this and hopefully with this
		exercise I'll highlight some of them and provide some
		solutions, being an engineer and all.`),

		P(`Throughout this exercise you will follow along a fiction
		story of an enterprise developing a space ship control
		system. In parts I'll use dialog form between stakeholders and
		engineers to highlight the iterative process required to
		produce easily digested specifications and requirements for
		software developers in particular. The specification can and
		is often a base which agreements are founded upon, so all
		stakeholders should be able to digest it easily, not only
		developers.`),

		P(`In this fictive process the customer is an internal
		departement, the ones building the space ship, John is their
		tech lead. They talk to engineers responsible for the
		software in our story it's Jane.`),

		John(`Hello, Jane! ready to start working on the control
		system?`),

		Jane(`Good morning, John! ready as can be, let's sit down.`),
		//
	)
}

func NewBeachStory() *Element {
	return Section(
		H2("To the beach"),

		P(`Through the navigation system people can plot a course or
		manually steer a ship.  People depend on its accuracy and
		automation to safely navigate through space.`),

		H3(`Plot new course`),

		P(`Standing at the bridge, the captain asks for the closest
		viable planets for some time at the beach. Selects the one
		with the nicest beaches and tells the system to plot the
		course. The plot details show that the route is through
		uncharted space. The captain selects another of the viable
		planets and tells the system to plot the course again. Once
		satisfied, he tells the system to engage.`),

		P(`The journey is estimated to five days. On the second day
		however an interference is detected in space and the ship
		adapts the course accordingly. The captain is notified through
		his personal communicator of the changes.`),

		H3(`Manual control`),

		P(`Once the ship enters the planets atmosphere one of the crew
	 	members on the bridge tells the system to let him manually
	 	steer the ship. He wants to find a suitable spot on the
	 	crowded beach, before letting the passengers leave the
	 	ship.`),
		//
	)
}

func NewElicitedFeatures(n *Hn) *Element {
	return Article(
		n.H1("Elicited features"),

		P(`These features have been elicitated from the navigation story`),

		n.H2("voice control"),
		n.H2("show route details"),
		n.H2("find destination"),
	)
}

func NewNavigationSystemSpec(n *Hn) *Element {

	return Article(
		n.H1("Navigation system"),

		Em(`Purpose; provide safe travel through space.`),

		P(`Through the navigation system people can plot a course or
        manually steer a ship.  People depend on its accuracy and
        automation to safely navigate through space.`),

		//
	)
}

func John(el ...interface{}) *Element {
	return Div("&#8213; John: ").With(el...)
}

func Jane(el ...interface{}) *Element {
	return Div("&#8213; Jane: ").With(el...)
}

func newPage(content ...interface{}) *Page {

	return NewPage(

		Html(
			Head(
				Meta(Charset("utf-8")),
				Style(theme()),
				Script(
					// to prevent Firefox FOUC, this must be here
					// https://stackoverflow.com/questions/21147149
					"let FF_FOUC_FIX;",
				),
			),
			Body(
				content...,
			),
		),
	)
}
